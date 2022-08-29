/*
 * @Author: jf-011101 2838264218@qq.com
 * @Date: 2022-08-02 14:03:24
 * @LastEditors: jf-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-19 22:04:41
 * @FilePath: \dytt\cmd\comment\main.go
 * @Description: Comment RPC server side initialization
 */

package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	zipkin "github.com/openzipkin/zipkin-go"
	zipkingrpc "github.com/openzipkin/zipkin-go/middleware/grpc"
	logreporter "github.com/openzipkin/zipkin-go/reporter/log"
	"google.golang.org/grpc"

	"github.com/jf-011101/dytt/dal"
	comment_pb "github.com/jf-011101/dytt/grpc_gen/comment"
	"github.com/jf-011101/dytt/internal/comment"
	"github.com/jf-011101/dytt/internal/pkg/discovery"
	"github.com/jf-011101/dytt/internal/pkg/gtls"
	"github.com/jf-011101/dytt/internal/pkg/ilog"
	my_grpc_middleware "github.com/jf-011101/dytt/internal/pkg/middleware/grpc"
	"github.com/jf-011101/dytt/internal/pkg/ttviper"
)

var (
	Config               = ttviper.ConfigInit("TIKTOK_COMMENT", "commentConfig")
	ServiceName          = Config.Viper.GetString("Server.Name")
	ServiceAddr          = fmt.Sprintf("%s:%d", Config.Viper.GetString("Server.Address"), Config.Viper.GetInt("Server.Port"))
	EtcdAddress          = fmt.Sprintf("%s:%d", Config.Viper.GetString("Etcd.Address"), Config.Viper.GetInt("Etcd.Port"))
	CertFile             = Config.Viper.GetString("TLS.CertFileLocalAddr")
	KeyFile              = Config.Viper.GetString("TLS.KeyFileLocalAddr")
	ZIPKIN_NAME          = Config.Viper.GetString("ZIPKIN.name")
	ZIPKIN_HTTP_ENDPOINT = Config.Viper.GetString("ZIPKIN.endpoint")
)

func Init() {
	dal.Init()
}

func main() {
	Init()

	etcdRegister := discovery.NewRegister([]string{EtcdAddress}, ilog.New(ilog.WithLevel(ilog.InfoLevel),
		ilog.WithFormatter(&ilog.JsonFormatter{IgnoreBasicFields: false}),
	))
	defer etcdRegister.Stop()
	commentNode := discovery.Server{
		Name: ServiceName,
		Addr: ServiceAddr,
	}
	if _, err := etcdRegister.Register(commentNode, 10); err != nil {
		ilog.Fatalf("register comment server failed, err: %v", err)
	}

	tlsServer := gtls.Server{
		CertFile: CertFile,
		KeyFile:  KeyFile,
	}
	c, err := tlsServer.GetTLSCredentials()
	if err != nil {
		ilog.Fatalf("tlsServer.GetTLSCredentials err: %v", err)
	}

	// set up a span reporter
	reporter := logreporter.NewReporter(log.New(os.Stderr, "", log.LstdFlags))
	defer reporter.Close()
	// create our local service endpoint
	endpoint, err := zipkin.NewEndpoint(ZIPKIN_NAME, ZIPKIN_HTTP_ENDPOINT)
	if err != nil {
		ilog.Fatalf("unable to create local endpoint: %+v\n", err)
	}
	// initialize our tracer
	tracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		ilog.Fatalf("unable to create tracer: %+v\n", err)
	}
	s := grpc.NewServer(grpc.Creds(c), grpc.StatsHandler(zipkingrpc.NewServerHandler(tracer)), grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
		grpc_auth.StreamServerInterceptor(my_grpc_middleware.AuthInterceptor),
		grpc_zap.StreamServerInterceptor(my_grpc_middleware.ZapInterceptor()),
		grpc_recovery.StreamServerInterceptor(my_grpc_middleware.RecoveryInterceptor()),
	)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_auth.UnaryServerInterceptor(my_grpc_middleware.AuthInterceptor),
			grpc_zap.UnaryServerInterceptor(my_grpc_middleware.ZapInterceptor()),
			grpc_recovery.UnaryServerInterceptor(my_grpc_middleware.RecoveryInterceptor()),
		)),
	)

	comment_pb.RegisterCommentSrvServer(s, &comment.CommentSrvImpl{})

	lis, err := net.Listen("tcp", ServiceAddr)
	if err != nil {
		ilog.Fatalf("failed to listen: %v", err)
	}
	go func() {
		if err := s.Serve(lis); err != nil {
			ilog.Fatalf("%s stopped with error: %v", ServiceName, err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ilog.Info("Shutting down comment rpc server...")

	// Delete node and revoke the given lease
	etcdRegister.Stop()

	ilog.Info("Server exiting")
}

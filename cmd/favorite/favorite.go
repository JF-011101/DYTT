/*
 * @Author: jf-011101 2838264218@qq.com
 * @Date: 2022-08-02 14:03:24
 * @LastEditors: jf-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-19 22:09:57
 * @FilePath: \dytt\cmd\favorite\main.go
 * @Description: Favorite RPC server initialization
 */

package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	zipkingrpc "github.com/openzipkin/zipkin-go/middleware/grpc"
	"google.golang.org/grpc"

	"github.com/jf-011101/dytt/dal"
	favorite_pb "github.com/jf-011101/dytt/grpc_gen/favorite"
	"github.com/jf-011101/dytt/internal/favorite"
	"github.com/jf-011101/dytt/internal/pkg/discovery"
	"github.com/jf-011101/dytt/internal/pkg/gtls"
	"github.com/jf-011101/dytt/internal/pkg/ilog"
	my_grpc_middleware "github.com/jf-011101/dytt/internal/pkg/middleware/grpc"
	"github.com/jf-011101/dytt/internal/pkg/tracing"
	"github.com/jf-011101/dytt/internal/pkg/ttviper"
)

var (
	Config          = ttviper.ConfigInit("TIKTOK_FAVORITE", "favoriteConfig")
	ServiceName     = Config.Viper.GetString("Server.Name")
	ServiceAddr     = fmt.Sprintf("%s:%d", Config.Viper.GetString("Server.Address"), Config.Viper.GetInt("Server.Port"))
	EtcdAddress     = fmt.Sprintf("%s:%d", Config.Viper.GetString("Etcd.Address"), Config.Viper.GetInt("Etcd.Port"))
	CertFile        = Config.Viper.GetString("TLS.CertFileLocalAddr")
	KeyFile         = Config.Viper.GetString("TLS.KeyFileLocalAddr")
	ZIPKIN_SRV_NAME = Config.Viper.GetString("ZIPKIN.SrvName")
	ZIPKIN_URL      = Config.Viper.GetString("ZIPKIN.Url")
	ZIPKIN_PORT     = Config.Viper.GetString("ZIPKIN.Port")
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
	favoriteNode := discovery.Server{
		Name: ServiceName,
		Addr: ServiceAddr,
	}

	tlsServer := gtls.Server{
		CertFile: CertFile,
		KeyFile:  KeyFile,
	}
	c, err := tlsServer.GetTLSCredentials()
	if err != nil {
		ilog.Fatalf("tlsServer.GetTLSCredentials err: %v", err)
	}

	tracer, _, err := tracing.NewZipkinTracer(ZIPKIN_URL, ZIPKIN_SRV_NAME, ZIPKIN_PORT)

	if err != nil {
		ilog.Fatalf("unable to create zipkin tracer: %+v\n", err)
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

	favorite_pb.RegisterFavoriteSrvServer(s, &favorite.FavoriteSrvImpl{})

	lis, err := net.Listen("tcp", ServiceAddr)
	if err != nil {
		ilog.Fatalf("failed to listen: %v", err)
	}

	if _, err := etcdRegister.Register(favoriteNode, 10); err != nil {
		ilog.Fatalf("register comment server failed, err: %v", err)
	}
	go func() {
		if err := s.Serve(lis); err != nil {
			ilog.Fatalf("%s stopped with error: %v", ServiceName, err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ilog.Info("Shutting down favorite rpc server...")

	// Delete node and revoke the given lease
	etcdRegister.Stop()

	ilog.Info("Server exiting")
}

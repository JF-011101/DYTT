/*
 * @Author: JF-011101 2838264218@qq.com
 * @Date: 2022-07-02 14:03:25
 * @LastEditors: JF-011101 2838264218@qq.com
 * @LastEditTime: 2022-07-21 11:36:55
 * @FilePath: \DYTT\cmd\user\main.go
 * @Description: User RPC server side initialization
 */

package main

import (
	"context"
	"fmt"
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"moul.io/zapgorm2"

	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/jf-011101/dytt/cmd/user/command"
	"github.com/jf-011101/dytt/dal"
	user "github.com/jf-011101/dytt/kitex_gen/user/usersrv"
	"github.com/jf-011101/dytt/pkg/dlog"
	"github.com/jf-011101/dytt/pkg/jwt"
	"github.com/jf-011101/dytt/pkg/middleware"
	"github.com/jf-011101/dytt/pkg/ttviper"
	etcd "github.com/jf-011101/registry-etcd"

	"github.com/cloudwego/kitex/server"

	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

var (
	Config       = ttviper.ConfigInit("TIKTOK_USER", "userConfig")
	ServiceName  = Config.Viper.GetString("Server.Name")
	ServiceAddr  = fmt.Sprintf("%s:%d", Config.Viper.GetString("Server.Address"), Config.Viper.GetInt("Server.Port"))
	EtcdAddress  = fmt.Sprintf("%s:%d", Config.Viper.GetString("Etcd.Address"), Config.Viper.GetInt("Etcd.Port"))
	Jwt          *jwt.JWT
	Argon2Config *command.Argon2Params
)

// User RPC Server 端配置初始化
func Init() {
	dal.Init()
	Jwt = jwt.NewJWT([]byte(Config.Viper.GetString("JWT.signingKey")))
	Argon2Config = &command.Argon2Params{
		Memory:      Config.Viper.GetUint32("Server.Argon2ID.Memory"),
		Iterations:  Config.Viper.GetUint32("Server.Argon2ID.Iterations"),
		Parallelism: uint8(Config.Viper.GetUint("Server.Argon2ID.Parallelism")),
		SaltLength:  Config.Viper.GetUint32("Server.Argon2ID.SaltLength"),
		KeyLength:   Config.Viper.GetUint32("Server.Argon2ID.KeyLength"),
	}
}

// User RPC Server 端运行
func main() {
	var logger dlog.ZapLogger = dlog.ZapLogger{
		Level: klog.LevelInfo,
	}
	zaplogger := zapgorm2.New(dlog.InitLog())
	logger.SugaredLogger.Base = &zaplogger

	klog.SetLogger(&logger)

	defer logger.SugaredLogger.Base.ZapLogger.Sync()

	r, err := etcd.NewEtcdRegistry([]string{EtcdAddress})
	if err != nil {
		klog.Fatal(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", ServiceAddr)
	if err != nil {
		klog.Fatal(err)
	}

	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(ServiceName),
		provider.WithExportEndpoint("localhost:4317"),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	Init()

	svr := user.NewServer(new(UserSrvImpl),
		server.WithServiceAddr(addr),                                       // address
		server.WithMiddleware(middleware.CommonMiddleware),                 // middleware
		server.WithMiddleware(middleware.ServerMiddleware),                 // middleware
		server.WithRegistry(r),                                             // registry
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		server.WithSuite(tracing.NewServerSuite()),                         // trace
		// Please keep the same as provider.WithServiceName
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: ServiceName}))

	if err := svr.Run(); err != nil {
		klog.Fatalf("%s stopped with error:", ServiceName, err)
	}
}

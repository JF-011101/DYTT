/*
 * @Author: JF-011101 2838264218@qq.com
 * @Date: 2022-07-02 14:03:25
 * @LastEditors: JF-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-19 22:10:17
 * @FilePath: \dytt\cmd\feed\main.go
 * @Description: Feed RPC server initialization
 */

package main

import (
	"fmt"
	"net"

	"github.com/jf-011101/dytt/dal"
	"github.com/jf-011101/dytt/grpc_gen/feed"
	"github.com/jf-011101/dytt/pkg/discover"
	"github.com/jf-011101/dytt/pkg/ilog"
	"github.com/jf-011101/dytt/pkg/jwt"
	"github.com/jf-011101/dytt/pkg/ttviper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

var (
	Config      = ttviper.ConfigInit("TIKTOK_FEED", "feedConfig")
	ServiceName = Config.Viper.GetString("Server.Name")
	ServiceAddr = fmt.Sprintf("%s:%d", Config.Viper.GetString("Server.Address"), Config.Viper.GetInt("Server.Port"))
	EtcdAddress = fmt.Sprintf("%s:%d", Config.Viper.GetString("Etcd.Address"), Config.Viper.GetInt("Etcd.Port"))
	Jwt         *jwt.JWT
)

type server struct {
	feed.FeedSrvServer
}

// Feed RPC Server 端配置初始化
func Init() {
	dal.Init()
	Jwt = jwt.NewJWT([]byte(Config.Viper.GetString("JWT.signingKey")))
}

// Feed RPC Server 端运行
func main() {
	// etcd注册
	etcdRegister := discover.NewResolver([]string{EtcdAddress}, ilog.New())
	resolver.Register(etcdRegister)

	Init()

	lis, err := net.Listen("tcp", ServiceAddr)
	if err != nil {
		ilog.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	feed.RegisterFeedSrvServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		ilog.Fatalf("%s stopped with error:", ServiceName, err)
	}
}

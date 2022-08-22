/*
 * @Author: JF-011101 2838264218@qq.com
 * @Date: 2022-07-02 14:03:24
 * @LastEditors: JF-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-19 22:04:41
 * @FilePath: \dytt\cmd\comment\main.go
 * @Description: Comment RPC server side initialization
 */

package main

import (
	"fmt"
	"net"

	"github.com/jf-011101/dytt/dal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"

	"github.com/jf-011101/dytt/pkg/discover"
	"github.com/jf-011101/dytt/pkg/ilog"
	"github.com/jf-011101/dytt/pkg/jwt"
	"github.com/jf-011101/dytt/pkg/ttviper"

	"github.com/jf-011101/dytt/grpc_gen/comment"
)

var (
	Config      = ttviper.ConfigInit("TIKTOK_COMMENT", "commentConfig")
	ServiceName = Config.Viper.GetString("Server.Name")
	ServiceAddr = fmt.Sprintf("%s:%d", Config.Viper.GetString("Server.Address"), Config.Viper.GetInt("Server.Port"))
	EtcdAddress = fmt.Sprintf("%s:%d", Config.Viper.GetString("Etcd.Address"), Config.Viper.GetInt("Etcd.Port"))
	Jwt         *jwt.JWT
)


// Comment RPC Server 端配置初始化
func Init() {
	dal.Init()
	Jwt = jwt.NewJWT([]byte(Config.Viper.GetString("JWT.signingKey")))
}

// Comment RPC Server 端运行
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
	comment.RegisterCommentSrvServer(s, &CommentSrvImpl{})

	if err := s.Serve(lis); err != nil {
		ilog.Fatalf("%s stopped with error:", ServiceName, err)
	}
}

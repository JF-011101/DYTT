/*
 * @Author: JF-011101 2838264218@qq.com
 * @Date: 2022-07-02 14:03:25
 * @LastEditors: JF-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-19 22:16:17
 * @FilePath: \dytt\cmd\user\main.go
 * @Description: User RPC server side initialization
 */

package main

import (
	"fmt"
	"net"

	"github.com/jf-011101/dytt/cmd/user/command"
	"github.com/jf-011101/dytt/dal"
	"github.com/jf-011101/dytt/grpc_gen/user"
	"github.com/jf-011101/dytt/pkg/discover"
	"github.com/jf-011101/dytt/pkg/ilog"
	"github.com/jf-011101/dytt/pkg/jwt"
	"github.com/jf-011101/dytt/pkg/ttviper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
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
	// etcd注册
	etcdRegister := discover.NewResolver([]string{EtcdAddress}, ilog.New())
	resolver.Register(etcdRegister)

	Init()

	lis, err := net.Listen("tcp", ServiceAddr)
	if err != nil {
		ilog.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	user.RegisterUserSrvServer(s, &UserSrvImpl{})

	if err := s.Serve(lis); err != nil {
		ilog.Fatalf("%s stopped with error:", ServiceName, err)
	}
}

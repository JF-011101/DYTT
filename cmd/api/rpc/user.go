/*
 * @Author: JF-011101 2838264218@qq.com
 * @Date: 2022-07-02 14:03:24
 * @LastEditors: JF-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-19 21:20:22
 * @FilePath: \dytt\cmd\api\rpc\user.go
 * @Description: User RPC client initialization
 and related RPC communication operation definitions
*/

package rpc

import (
	"context"
	"fmt"
	"time"

	"github.com/jf-011101/dytt/grpc_gen/user"
	"github.com/jf-011101/dytt/pkg/discover"
	"github.com/jf-011101/dytt/pkg/errno"
	"github.com/jf-011101/dytt/pkg/ilog"
	"github.com/jf-011101/dytt/pkg/ttviper"
	"google.golang.org/grpc/resolver"
)

var userClient user.UserSrvClient

// User RPC 客户端初始化
func initUserRpc(Config *ttviper.Config) {
	EtcdAddress := fmt.Sprintf("%s:%d", Config.Viper.GetString("Etcd.Address"), Config.Viper.GetInt("Etcd.Port"))
	ServerAddress := fmt.Sprintf("%s:%d", Config.Viper.GetString("Server.Address"), Config.Viper.GetInt("Server.Port"))

	etcdRegister := discover.NewResolver([]string{EtcdAddress}, ilog.New())
	resolver.Register(etcdRegister)
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	// RPC 连接
	conn, err := RPCConnect(ctx, ServerAddress, etcdRegister)
	if err != nil {
		panic(err)
	}
	userClient = user.NewUserSrvClient(conn)
}

// 传递 注册操作 的上下文, 并获取 RPC Server 端的响应.
func Register(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	resp, err = userClient.Register(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}

// 传递 登录操作 的上下文, 并获取 RPC Server 端的响应.
func Login(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	resp, err = userClient.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}

// 传递 获取用户信息操作 的上下文, 并获取 RPC Server 端的响应.
func GetUserById(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	resp, err = userClient.GetUserById(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}

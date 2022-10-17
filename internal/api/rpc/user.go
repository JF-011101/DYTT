/*
 * @Author: jf-011101 2838264218@qq.com
 * @Date: 2022-08-02 14:03:24
 * @LastEditors: jf-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-19 21:24:01
 * @FilePath: \dytt\cmd\api\rpc\user.go
 * @Description: Usership RPC client initialization
 and related RPC communication operation definitions
*/

package rpc

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/resolver"

	"github.com/jf-011101/dytt/grpc_gen/user"
	"github.com/jf-011101/dytt/internal/pkg/discovery"
	"github.com/jf-011101/dytt/internal/pkg/gtls"
	"github.com/jf-011101/dytt/internal/pkg/ilog"
	"github.com/jf-011101/dytt/internal/pkg/tracing"
	"github.com/jf-011101/dytt/internal/pkg/ttviper"
	"github.com/jf-011101/dytt/pkg/errno"
)

var userClient user.UserSrvClient

func initUserRpc(Config *ttviper.Config) {
	ServerAddress := fmt.Sprintf("%s:%d", Config.Viper.GetString("Server.Address"), Config.Viper.GetInt("Server.Port"))
	ZIPKIN_CLI_NAME := Config.Viper.GetString("ZIPKIN.CliName")
	ZIPKIN_URL := Config.Viper.GetString("ZIPKIN.Url")
	ZIPKIN_PORT := Config.Viper.GetString("ZIPKIN.Port")

	// etcd register
	EtcdAddress := fmt.Sprintf("%s:%d", Config.Viper.GetString("Etcd.Address"), Config.Viper.GetInt("Etcd.Port"))
	EtcdResolver := discovery.NewResolver([]string{EtcdAddress}, ilog.New())
	resolver.Register(EtcdResolver)
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	fmt.Print("etcd rig over")

	// init tlsClient and token
	tlsClient := gtls.Client{
		ServerName: "dytt.user.com",
		CertFile:   "../../config/cert/user/dytt-user.pem",
	}
	c, err := tlsClient.GetTLSCredentials()
	if err != nil {
		ilog.Fatalf("tlsClient.GetTLSCredentials err: %v", err)
	}
	token := Token{
		Value: "bearer dytt.grpc.auth.token",
	}

	tracer, _, err := tracing.NewZipkinTracer(ZIPKIN_URL, ZIPKIN_CLI_NAME, ZIPKIN_PORT)

	if err != nil {
		ilog.Fatalf("unable to create zipkin tracer: %+v\n", err)
	}

	conn, err := RPCConnect(ctx, ServerAddress, &token, tracer, c)
	if err != nil {
		panic(err)
	}
	userClient = user.NewUserSrvClient(conn)
}

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

func Refresh(ctx context.Context, req *user.DouyinUserRefreshRequest) (resp *user.DouyinUserRefreshResponse, err error) {
	resp, err = userClient.Refresh(ctx, req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func QueryUser(ctx context.Context, req *user.DouyinUserQueryRequest) (resp *user.DouyinUserQueryResponse, err error) {
	resp, err = userClient.QueryUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}

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

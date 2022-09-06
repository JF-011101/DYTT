/*
 * @Author: jf-011101 2838264218@qq.com
 * @Date: 2022-08-02 14:03:24
 * @LastEditors: jf-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-19 21:31:06
 * @FilePath: \dytt\cmd\api\rpc\publish.go
 * @Description: Publish RPC client initialization
 and related RPC communication operation definitions
*/

package rpc

import (
	"context"
	"fmt"
	"time"

	"github.com/jf-011101/dytt/grpc_gen/publish"
	"github.com/jf-011101/dytt/internal/pkg/discovery"
	"github.com/jf-011101/dytt/internal/pkg/gtls"
	"github.com/jf-011101/dytt/internal/pkg/ilog"
	"github.com/jf-011101/dytt/internal/pkg/tracing"
	"github.com/jf-011101/dytt/internal/pkg/ttviper"
	"github.com/jf-011101/dytt/pkg/errno"

	"google.golang.org/grpc/resolver"
)

var publishClient publish.PublishSrvClient

func initPublishRpc(Config *ttviper.Config) {
	ServerAddress := fmt.Sprintf("%s:%d", Config.Viper.GetString("Server.Address"), Config.Viper.GetInt("Server.Port"))
	ZIPKIN_CLI_NAME := Config.Viper.GetString("ZIPKIN.CliName")
	ZIPKIN_URL := Config.Viper.GetString("ZIPKIN.Url")
	ZIPKIN_PORT := Config.Viper.GetString("ZIPKIN.Port")
	// init tlsClient and token
	tlsClient := gtls.Client{
		ServerName: "dytt.publish.com",
		CertFile:   "../../config/cert/publish/dytt-publish.pem",
	}
	c, err := tlsClient.GetTLSCredentials()

	if err != nil {
		ilog.Fatalf("tlsClient.GetTLSCredentials err: %v", err)
	}
	token := &Token{
		Value: "bearer dytt.grpc.auth.token",
	}
	// etcd register
	EtcdAddress := fmt.Sprintf("%s:%d", Config.Viper.GetString("Etcd.Address"), Config.Viper.GetInt("Etcd.Port"))
	EtcdRegister := discovery.NewResolver([]string{EtcdAddress}, ilog.New())
	resolver.Register(EtcdRegister)
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	tracer, _, err := tracing.NewZipkinTracer(ZIPKIN_URL, ZIPKIN_CLI_NAME, ZIPKIN_PORT)

	if err != nil {
		ilog.Fatalf("unable to create zipkin tracer: %+v\n", err)
	}

	conn, err := RPCConnect(ctx, ServerAddress, token, tracer, c)
	if err != nil {
		panic(err)
	}
	publishClient = publish.NewPublishSrvClient(conn)
}

func PublishAction(ctx context.Context, req *publish.DouyinPublishActionRequest) (resp *publish.DouyinPublishActionResponse, err error) {
	resp, err = publishClient.PublishAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}

func PublishList(ctx context.Context, req *publish.DouyinPublishListRequest) (resp *publish.DouyinPublishListResponse, err error) {
	resp, err = publishClient.PublishList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}

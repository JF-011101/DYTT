/*
 * @Author: jf-011101 2838264218@qq.com
 * @Date: 2022-08-02 14:03:24
 * @LastEditors: jf-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-19 21:48:38
 * @FilePath: \dytt\cmd\api\rpc\comment.go
 * @Description: Comment RPC client initialization
 and related RPC communication operation definitions
*/

package rpc

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/resolver"

	"github.com/jf-011101/dytt/grpc_gen/comment"
	"github.com/jf-011101/dytt/internal/pkg/discovery"
	"github.com/jf-011101/dytt/internal/pkg/gtls"
	"github.com/jf-011101/dytt/internal/pkg/ilog"
	"github.com/jf-011101/dytt/internal/pkg/tracing"
	"github.com/jf-011101/dytt/internal/pkg/ttviper"
	"github.com/jf-011101/dytt/pkg/errno"
)

var commentClient comment.CommentSrvClient

func initCommentRpc(Config *ttviper.Config) {
	ServerAddress := fmt.Sprintf("%s:%d", Config.Viper.GetString("Server.Address"), Config.Viper.GetInt("Server.Port"))
	ZIPKIN_CLI_NAME := Config.Viper.GetString("ZIPKIN.CliName")
	ZIPKIN_URL := Config.Viper.GetString("ZIPKIN.Url")
	ZIPKIN_PORT := Config.Viper.GetString("ZIPKIN.Port")

	// init tlsClient and token
	tlsClient := gtls.Client{
		ServerName: "dytt.comment.com",
		CertFile:   "../../config/cert/api/dytt-api.pem",
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
	EtcdResolver := discovery.NewResolver([]string{EtcdAddress}, ilog.New())
	resolver.Register(EtcdResolver)

	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	tracer, _, err := tracing.NewZipkinTracer(ZIPKIN_URL, ZIPKIN_CLI_NAME, ZIPKIN_PORT)
	if err != nil {
		ilog.Fatalf("unable to create zipkin tracer: %+v\n", err)
	}

	conn, err := RPCConnect(ctx, ServerAddress, token, tracer, c)
	if err != nil {
		panic(err)
	}

	commentClient = comment.NewCommentSrvClient(conn)

}

func CommentAction(ctx context.Context, req *comment.DouyinCommentActionRequest) (resp *comment.DouyinCommentActionResponse, err error) {
	resp, err = commentClient.CommentAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}

func CommentList(ctx context.Context, req *comment.DouyinCommentListRequest) (resp *comment.DouyinCommentListResponse, err error) {
	resp, err = commentClient.CommentList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}

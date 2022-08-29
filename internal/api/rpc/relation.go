/*
 * @Author: jf-011101 2838264218@qq.com
 * @Date: 2022-08-02 14:03:24
 * @LastEditors: jf-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-19 21:24:01
 * @FilePath: \dytt\cmd\api\rpc\relation.go
 * @Description: Relationship RPC client initialization
 and related RPC communication operation definitions
*/

package rpc

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jf-011101/dytt/grpc_gen/relation"
	"github.com/jf-011101/dytt/internal/pkg/discovery"
	"github.com/jf-011101/dytt/internal/pkg/gtls"
	"github.com/jf-011101/dytt/internal/pkg/ilog"
	"github.com/jf-011101/dytt/internal/pkg/ttviper"
	"github.com/jf-011101/dytt/pkg/errno"
	zipkin "github.com/openzipkin/zipkin-go"
	logreporter "github.com/openzipkin/zipkin-go/reporter/log"

	"google.golang.org/grpc/resolver"
)

var relationClient relation.RelationSrvClient

func initRelationRpc(Config *ttviper.Config) {
	ServerAddress := fmt.Sprintf("%s:%d", Config.Viper.GetString("Server.Address"), Config.Viper.GetInt("Server.Port"))
	ZIPKIN_NAME := Config.Viper.GetString("ZIPKIN.name")
	ZIPKIN_HTTP_ENDPOINT := Config.Viper.GetString("ZIPKIN.endpoint")
	// init tlsClient and token
	tlsClient := gtls.Client{
		ServerName: "dytt.relation.com",
		CertFile:   "../../config/cert/relation/dytt-relation.pem",
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

	// open-tracing
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

	conn, err := RPCConnect(ctx, ServerAddress, token, tracer, c)
	if err != nil {
		panic(err)
	}
	relationClient = relation.NewRelationSrvClient(conn)
}

func RelationAction(ctx context.Context, req *relation.DouyinRelationActionRequest) (resp *relation.DouyinRelationActionResponse, err error) {
	resp, err = relationClient.RelationAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}

func RelationFollowList(ctx context.Context, req *relation.DouyinRelationFollowListRequest) (resp *relation.DouyinRelationFollowListResponse, err error) {
	resp, err = relationClient.RelationFollowList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}

func RelationFollowerList(ctx context.Context, req *relation.DouyinRelationFollowerListRequest) (resp *relation.DouyinRelationFollowerListResponse, err error) {
	resp, err = relationClient.RelationFollowerList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}

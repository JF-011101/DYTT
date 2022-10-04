/*
 * @Author: jf-011101 2838264218@qq.com
 * @Date: 2022-08-02 14:03:24
 * @LastEditors: jf-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-19 21:56:40
 * @FilePath: \dytt\cmd\api\rpc\favorite.go
 * @Description: Favorite RPC client initialization
 and related RPC communication operation definitions
*/

package rpc

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/resolver"

	"github.com/jf-011101/dytt/grpc_gen/favorite"
	"github.com/jf-011101/dytt/internal/pkg/discovery"
	"github.com/jf-011101/dytt/internal/pkg/gtls"
	"github.com/jf-011101/dytt/internal/pkg/ilog"
	"github.com/jf-011101/dytt/internal/pkg/tracing"
	"github.com/jf-011101/dytt/internal/pkg/ttviper"
	"github.com/jf-011101/dytt/pkg/errno"
)

var favoriteClient favorite.FavoriteSrvClient

func initFavoriteRpc(Config *ttviper.Config) {
	ServerAddress := fmt.Sprintf("%s:%d", Config.Viper.GetString("Server.Address"), Config.Viper.GetInt("Server.Port"))
	ZIPKIN_CLI_NAME := Config.Viper.GetString("ZIPKIN.CliName")
	ZIPKIN_URL := Config.Viper.GetString("ZIPKIN.Url")
	ZIPKIN_PORT := Config.Viper.GetString("ZIPKIN.Port")

	// init tlsClient and token
	tlsClient := gtls.Client{
		ServerName: "dytt.favorite.com",
		CertFile:   "../../config/cert/favorite/dytt-favorite.pem",
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

	favoriteClient = favorite.NewFavoriteSrvClient(conn)
}

func FavoriteAction(ctx context.Context, req *favorite.DouyinFavoriteActionRequest) (resp *favorite.DouyinFavoriteActionResponse, err error) {
	resp, err = favoriteClient.FavoriteAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}

func FavoriteList(ctx context.Context, req *favorite.DouyinFavoriteListRequest) (resp *favorite.DouyinFavoriteListResponse, err error) {
	resp, err = favoriteClient.FavoriteList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}

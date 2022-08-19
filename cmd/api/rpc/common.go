/*
 * @Author: JF-011101 2838264218@qq.com
 * @Date: 2022-08-19 19:08:24
 * @LastEditors: JF-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-19 19:30:48
 * @FilePath: \dytt\cmd\api\rpc\common.go
 * @Description: functions and args
 */
package rpc

import (
	"context"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/jf-011101/dytt/pkg/discover"
)

func RPCConnect(ctx context.Context, serviceAddr string, etcdRegister *discover.Resolver) (conn *grpc.ClientConn, err error) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	}
	conn, err = grpc.DialContext(ctx, serviceAddr, opts...)
	return
}

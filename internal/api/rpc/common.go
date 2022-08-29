/*
 * @Author: jf-011101 2838264218@qq.com
 * @Date: 2022-08-19 19:08:24
 * @LastEditors: jf-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-19 19:30:48
 * @FilePath: \dytt\cmd\api\rpc\common.go
 * @Description: RPC client connect and custom authauthentication of grpc
 */
package rpc

import (
	"context"

	"github.com/jf-011101/dytt/internal/pkg/ilog"
	zipkin "github.com/openzipkin/zipkin-go"
	zipkingrpc "github.com/openzipkin/zipkin-go/middleware/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Token struct {
	Value string
}

const headerAuthorize string = "authorization"

// GetRequestMetadata 获取当前请求认证所需的元数据
func (t *Token) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{headerAuthorize: t.Value}, nil
}

// RequireTransportSecurity 是否需要基于 TLS 认证进行安全传输
func (t *Token) RequireTransportSecurity() bool {
	return true
}

func RPCConnect(ctx context.Context, serviceAddr string, token *Token, tracer *zipkin.Tracer, c credentials.TransportCredentials) (conn *grpc.ClientConn, err error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(c),
		grpc.WithPerRPCCredentials(token),
		grpc.WithStatsHandler(zipkingrpc.NewClientHandler(tracer)),
	}
	conn, err = grpc.DialContext(ctx, serviceAddr, opts...)
	if err != nil {
		ilog.Fatalf("net.Connect err: %v", err)
	}
	return
}

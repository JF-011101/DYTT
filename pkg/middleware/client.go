/*
 * @Author: JF-011101 2838264218@qq.com
 * @Date: 2022-07-02 14:03:26
 * @LastEditors: JF-011101 2838264218@qq.com
 * @LastEditTime: 2022-07-21 11:10:58
 * @FilePath: \DYTT\pkg\middleware\client.go
 * @Description: RPC Client Middleware
 */

package middleware

import (
	"context"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

var _ endpoint.Middleware = ClientMiddleware

// ClientMiddleware client middleware print server address 、rpc timeout and connection timeout
func ClientMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		ri := rpcinfo.GetRPCInfo(ctx)
		// get server information
		klog.Infof("server address: %v, rpc timeout: %v, readwrite timeout: %v", ri.To().Address(), ri.Config().RPCTimeout(), ri.Config().ConnectTimeout())
		if err = next(ctx, req, resp); err != nil {
			return err
		}
		return nil
	}
}

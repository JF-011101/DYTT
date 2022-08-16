/*
 * @Author: JF-011101 2838264218@qq.com
 * @Date: 2022-07-02 14:03:26
 * @LastEditors: JF-011101 2838264218@qq.com
 * @LastEditTime: 2022-07-21 11:11:07
 * @FilePath: \DYTT\pkg\middleware\common.go
 * @Description: RPC Common Middleware
 */

package middleware

import (
	"context"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/jf-011101/dytt/pkg/dlog"
	"moul.io/zapgorm2"
)

var _ endpoint.Middleware = CommonMiddleware

func init() {
	var logger dlog.ZapLogger = dlog.ZapLogger{
		Level: klog.LevelInfo,
	}

	zaplogger := zapgorm2.New(dlog.InitLog())
	logger.SugaredLogger.Base = &zaplogger

	klog.SetLogger(&logger)
}

// CommonMiddleware common middleware print some rpc info、real request and real response
func CommonMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		ri := rpcinfo.GetRPCInfo(ctx)
		// get real request
		klog.Debugf("real request: %+v", req)
		// get remote service information
		klog.Debugf("remote service name: %s, remote method: %s", ri.To().ServiceName(), ri.To().Method())
		if err = next(ctx, req, resp); err != nil {
			return err
		}
		// get real response
		klog.Infof("real response: %+v\n", resp)
		return nil
	}
}

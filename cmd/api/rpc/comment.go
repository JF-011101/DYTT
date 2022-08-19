/*
 * @Author: JF-011101 2838264218@qq.com
 * @Date: 2022-07-02 14:03:24
 * @LastEditors: JF-011101 2838264218@qq.com
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

	// "github.com/cloudwego/kitex/client"
	// "github.com/cloudwego/kitex/pkg/retry"
	// "github.com/cloudwego/kitex/pkg/rpcinfo"

	"github.com/jf-011101/dytt/grpc_gen/comment"
	"github.com/jf-011101/dytt/pkg/discover"
	"github.com/jf-011101/dytt/pkg/errno"
	"github.com/jf-011101/dytt/pkg/ilog"
	"github.com/jf-011101/dytt/pkg/ttviper"
	"google.golang.org/grpc/resolver"
)

var commentClient comment.CommentSrvClient

// Comment RPC 客户端初始化
func initCommentRpc(Config *ttviper.Config) {
	EtcdAddress := fmt.Sprintf("%s:%d", Config.Viper.GetString("Etcd.Address"), Config.Viper.GetInt("Etcd.Port"))
	ServerAddress := fmt.Sprintf("%s:%d", Config.Viper.GetString("Server.Address"), Config.Viper.GetInt("Server.Port"))

	// etcd注册
	etcdRegister := discover.NewResolver([]string{EtcdAddress}, ilog.New())
	resolver.Register(etcdRegister)
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	// RPC 连接
	conn, err := RPCConnect(ctx, ServerAddress, etcdRegister)
	if err != nil {
		panic(err)
	}

	commentClient = comment.NewCommentSrvClient(conn)

}

// 传递 评论操作 的上下文, 并获取 RPC Server 端的响应.
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

// 传递 获取评论列表操作 的上下文, 并获取 RPC Server 端的响应.
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

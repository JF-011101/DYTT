/*
 * @Author: jf-011101 2838264218@qq.com
 * @Date: 2022-06-12 14:03:24
 * @LastEditors: jf-011101 2838264218@qq.com
 * @LastEditTime: 2022-06-18 10:59:08
 * @FilePath: \dytt\cmd\api\handlers\comment.go
 * @Description: define Comment API's handler, pass context to rpc client and get response
 */
package handlers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/jf-011101/dytt/dal/pack"
	"github.com/jf-011101/dytt/grpc_gen/comment"
	"github.com/jf-011101/dytt/internal/api/rpc"
	"github.com/jf-011101/dytt/pkg/errno"
)

func CommentAction(c *gin.Context) {
	var paramVar CommentActionParam
	token := c.PostForm("token")
	video_id := c.PostForm("video_id")
	action_type := c.PostForm("action_type")

	vid, err := strconv.Atoi(video_id)
	if err != nil {
		SendResponse(c, pack.BuildCommentActionResp(errno.ErrBind))
		return
	}
	act, err := strconv.Atoi(action_type)
	if err != nil {
		SendResponse(c, pack.BuildCommentActionResp(errno.ErrBind))
		return
	}

	paramVar.Token = token
	paramVar.VideoId = int64(vid)
	paramVar.ActionType = int32(act)

	rpcReq := comment.DouyinCommentActionRequest{
		VideoId:    paramVar.VideoId,
		Token:      paramVar.Token,
		ActionType: paramVar.ActionType,
	}

	if act == 1 {
		comment_text := c.PostForm("comment_text")
		rpcReq.CommentText = &comment_text
	} else {
		comment_id := c.PostForm("comment_id")
		cid, err := strconv.Atoi(comment_id)
		if err != nil {
			SendResponse(c, pack.BuildCommentActionResp(errno.ErrBind))
			return
		}
		cid64 := int64(cid)
		rpcReq.CommentId = &cid64
	}

	resp, err := rpc.CommentAction(context.Background(), &rpcReq)
	if err != nil {
		SendResponse(c, pack.BuildCommentActionResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}

func CommentList(c *gin.Context) {
	var paramVar CommentListParam
	videoid, err := strconv.Atoi(c.Query("video_id"))
	if err != nil {
		SendResponse(c, pack.BuildCommentListResp(errno.ErrBind))
		return
	}
	paramVar.VideoId = int64(videoid)
	paramVar.Token = c.Query("token")

	if len(paramVar.Token) == 0 || paramVar.VideoId < 0 {
		SendResponse(c, pack.BuildCommentListResp(errno.ErrBind))
		return
	}

	resp, err := rpc.CommentList(context.Background(), &comment.DouyinCommentListRequest{
		VideoId: paramVar.VideoId,
		Token:   paramVar.Token,
	})
	if err != nil {
		SendResponse(c, pack.BuildCommentListResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}

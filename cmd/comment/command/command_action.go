/*
 * @Author: JF-011101 2838264218@qq.com
 * @Date: 2022-07-02 14:03:24
 * @LastEditors: JF-011101 2838264218@qq.com
 * @LastEditTime: 2022-07-21 11:28:42
 * @FilePath: \dytt\cmd\comment\command\command_action.go
 * @Description: Comment's operation business logic
 */

package command

import (
	"context"

	"github.com/jf-011101/dytt/grpc_gen/comment"
	"github.com/jf-011101/dytt/pkg/errno"

	"github.com/jf-011101/dytt/dal/db"
)

type CommentActionService struct {
	ctx context.Context
}

// NewCommentActionService new CommentActionService
func NewCommentActionService(ctx context.Context) *CommentActionService {
	return &CommentActionService{ctx: ctx}
}

// CommentActionService action comment.
func (s *CommentActionService) CommentAction(req *comment.DouyinCommentActionRequest) error {
	// 1-评论
	if req.ActionType == 1 {
		return db.NewComment(s.ctx, &db.Comment{
			UserID:  int(req.UserId),
			VideoID: int(req.VideoId),
			Content: *req.CommentText,
		})
	}
	// 2-删除评论
	if req.ActionType == 2 {
		return db.DelComment(s.ctx, *req.CommentId, req.VideoId)
	}
	return errno.ErrBind
}

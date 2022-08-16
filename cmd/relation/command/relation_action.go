/*
 * @Author: JF-011101 2838264218@qq.com
 * @Date: 2022-07-02 14:03:25
 * @LastEditors: JF-011101 2838264218@qq.com
 * @LastEditTime: 2022-07-21 11:34:14
 * @FilePath: \DYTT\cmd\relation\command\relation_action.go
 * @Description: Focus on user operation business logic
 */

package command

import (
	"context"

	"github.com/jf-011101/dytt/kitex_gen/relation"
	"github.com/jf-011101/dytt/pkg/errno"

	"github.com/jf-011101/dytt/dal/db"
)

type RelationActionService struct {
	ctx context.Context
}

// NewRelationActionService new RelationActionService
func NewRelationActionService(ctx context.Context) *RelationActionService {
	return &RelationActionService{ctx: ctx}
}

// RelationAction action favorite.
func (s *RelationActionService) RelationAction(req *relation.DouyinRelationActionRequest) error {
	// 1-关注
	if req.ActionType == 1 {
		return db.NewRelation(s.ctx, req.UserId, req.ToUserId)
	}
	// 2-取消关注
	if req.ActionType == 2 {
		return db.DisRelation(s.ctx, req.UserId, req.ToUserId)
	}
	return errno.ErrBind
}

/*
 * @Author: jf-011101 2838264218@qq.com
 * @Date: 2022-08-02 14:03:25
 * @LastEditors: jf-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-21 11:34:37
 * @FilePath: \dytt\cmd\relation\command\relation_list.go
 * @Description: Business logic for obtaining attention list operation
 */

package command

import (
	"context"

	"github.com/jf-011101/dytt/dal/pack"
	"github.com/jf-011101/dytt/grpc_gen/relation"
	"github.com/jf-011101/dytt/grpc_gen/user"

	"github.com/jf-011101/dytt/dal/db"
)

type FollowingListService struct {
	ctx context.Context
}

// NewFollowingListService creates a new FollowingListService
func NewFollowingListService(ctx context.Context) *FollowingListService {
	return &FollowingListService{
		ctx: ctx,
	}
}

// FollowingList returns the following lists
func (s *FollowingListService) FollowingList(req *relation.DouyinRelationFollowListRequest, fromID int64) ([]*user.User, error) {
	FollowingUser, err := db.FollowingList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return pack.FollowingList(s.ctx, FollowingUser, fromID)
}

type FollowerListService struct {
	ctx context.Context
}

// NewFollowerListService creates a new FollowerListService
func NewFollowerListService(ctx context.Context) *FollowerListService {
	return &FollowerListService{
		ctx: ctx,
	}
}

// FollowerList returns the Follower Lists
func (s *FollowerListService) FollowerList(req *relation.DouyinRelationFollowerListRequest, fromID int64) ([]*user.User, error) {
	FollowerUser, err := db.FollowerList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return pack.FollowerList(s.ctx, FollowerUser, fromID)
}

/*
 * @Author: JF-011101 2838264218@qq.com
 * @Date: 2022-07-02 14:03:25
 * @LastEditors: JF-011101 2838264218@qq.com
 * @LastEditTime: 2022-07-21 11:34:55
 * @FilePath: \DYTT\cmd\relation\handler.go
 * @Description: Define the relevant interfaces on the relationship RPC server side
 */

package main

import (
	"context"

	"github.com/jf-011101/dytt/cmd/relation/command"
	"github.com/jf-011101/dytt/dal/pack"
	"github.com/jf-011101/dytt/kitex_gen/relation"
	"github.com/jf-011101/dytt/pkg/errno"
)

// RelationSrvImpl implements the last service interface defined in the IDL.
type RelationSrvImpl struct{}

// RelationAction implements the RelationSrvImpl interface.
func (s *RelationSrvImpl) RelationAction(ctx context.Context, req *relation.DouyinRelationActionRequest) (resp *relation.DouyinRelationActionResponse, err error) {
	claim, err := Jwt.ParseToken(req.Token)
	if err != nil {
		resp = pack.BuildRelationActionResp(errno.ErrTokenInvalid)
		return resp, nil
	}

	if req.UserId == 0 || claim.Id != 0 {
		req.UserId = claim.Id
	}

	if req.ActionType < 1 || req.ActionType > 2 {
		resp = pack.BuildRelationActionResp(errno.ErrBind)
		return resp, nil
	}
	err = command.NewRelationActionService(ctx).RelationAction(req)
	if err != nil {
		resp = pack.BuildRelationActionResp(err)
		return resp, nil
	}
	resp = pack.BuildRelationActionResp(errno.Success)
	return resp, nil
}

// RelationFollowList implements the RelationSrvImpl interface.
func (s *RelationSrvImpl) RelationFollowList(ctx context.Context, req *relation.DouyinRelationFollowListRequest) (resp *relation.DouyinRelationFollowListResponse, err error) {
	claim, err := Jwt.ParseToken(req.Token)
	if err != nil {
		resp = pack.BuildFollowingListResp(errno.ErrTokenInvalid)
		return resp, nil
	}

	if req.UserId == 0 || claim.Id != 0 {
		req.UserId = claim.Id // 没有传入UserID，默认为自己
	}

	users, err := command.NewFollowingListService(ctx).FollowingList(req, claim.Id)
	if err != nil {
		resp = pack.BuildFollowingListResp(err)
		return resp, nil
	}

	resp = pack.BuildFollowingListResp(errno.Success)
	resp.UserList = users
	return resp, nil
}

// RelationFollowerList implements the RelationSrvImpl interface.
func (s *RelationSrvImpl) RelationFollowerList(ctx context.Context, req *relation.DouyinRelationFollowerListRequest) (resp *relation.DouyinRelationFollowerListResponse, err error) {
	claim, err := Jwt.ParseToken(req.Token)
	if err != nil {
		resp = pack.BuildFollowerListResp(errno.ErrTokenInvalid)
		return resp, nil
	}

	if req.UserId == 0 {
		req.UserId = claim.Id // 没有传入UserID，默认为自己
	}

	users, err := command.NewFollowerListService(ctx).FollowerList(req, claim.Id)
	if err != nil {
		resp = pack.BuildFollowerListResp(err)
		return resp, nil
	}

	resp = pack.BuildFollowerListResp(errno.Success)
	resp.UserList = users
	return resp, nil
}

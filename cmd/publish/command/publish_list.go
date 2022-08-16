/*
 * @Author: JF-011101 2838264218@qq.com
 * @Date: 2022-07-02 14:03:25
 * @LastEditors: JF-011101 2838264218@qq.com
 * @LastEditTime: 2022-07-21 11:33:33
 * @FilePath: \DYTT\cmd\publish\command\publish_list.go
 * @Description: Business logic for obtaining and publishing video list
 */

package command

import (
	"context"

	"github.com/jf-011101/dytt/dal/pack"
	"github.com/jf-011101/dytt/kitex_gen/feed"
	"github.com/jf-011101/dytt/kitex_gen/publish"

	"github.com/jf-011101/dytt/dal/db"
)

type PublishListService struct {
	ctx context.Context
}

// NewPublishListService new PublishListService
func NewPublishListService(ctx context.Context) *PublishListService {
	return &PublishListService{ctx: ctx}
}

// PublishList publish video.
func (s *PublishListService) PublishList(req *publish.DouyinPublishListRequest) (vs []*feed.Video, err error) {
	videos, err := db.PublishList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	vs, err = pack.Videos(s.ctx, videos, &req.UserId)
	if err != nil {
		return nil, err
	}

	return vs, nil
}

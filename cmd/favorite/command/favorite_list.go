/*
 * @Author: JF-011101 2838264218@qq.com
 * @Date: 2022-07-02 14:03:25
 * @LastEditors: JF-011101 2838264218@qq.com
 * @LastEditTime: 2022-07-21 11:31:01
 * @FilePath: \DYTT\cmd\favorite\command\favorite_list.go
 * @Description: Business logic for obtaining likes list
 */

package command

import (
	"context"

	"github.com/jf-011101/dytt/dal/pack"
	"github.com/jf-011101/dytt/kitex_gen/favorite"
	"github.com/jf-011101/dytt/kitex_gen/feed"

	"github.com/jf-011101/dytt/dal/db"
)

type FavoriteListService struct {
	ctx context.Context
}

// NewFavoriteListService creates a new FavoriteListService
func NewFavoriteListService(ctx context.Context) *FavoriteListService {
	return &FavoriteListService{
		ctx: ctx,
	}
}

// FavoriteList returns a Favorite List
func (s *FavoriteListService) FavoriteList(req *favorite.DouyinFavoriteListRequest) ([]*feed.Video, error) {
	FavoriteVideos, err := db.FavoriteList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	videos, err := pack.FavoriteVideos(s.ctx, FavoriteVideos, &req.UserId)
	if err != nil {
		return nil, err
	}
	return videos, nil
}

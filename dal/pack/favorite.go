/*
 * @Author: JF-011101 2838264218@qq.com
 * @Date: 2022-07-02 14:03:25
 * @LastEditors: JF-011101 2838264218@qq.com
 * @LastEditTime: 2022-07-21 11:39:49
 * @FilePath: \dytt\dal\pack\favorite.go
 * @Description: Encapsulate favoritevideos database data as RPC server-side response
 */

package pack

import (
	"context"

	"github.com/jf-011101/dytt/grpc_gen/feed"

	"github.com/jf-011101/dytt/dal/db"
)

// FavoriteVideos pack favoriteVideos info.
func FavoriteVideos(ctx context.Context, vs []db.Video, uid *int64) ([]*feed.Video, error) {
	videos := make([]*db.Video, 0)
	for _, v := range vs {
		videos = append(videos, &v)
	}

	packVideos, err := Videos(ctx, videos, uid)
	if err != nil {
		return nil, err
	}

	return packVideos, nil
}

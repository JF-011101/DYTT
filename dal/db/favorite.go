/*
 * @Author: jf-011101 2838264218@qq.com
 * @Date: 2022-08-02 14:03:25
 * @LastEditors: jf-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-21 11:06:32
 * @FilePath: \dytt\dal\db\favorite.go
 * @Description: Favorite database operation business logic
 */

package db

import (
	"context"

	"github.com/jf-011101/dytt/pkg/errno"
	"gorm.io/gorm"
)

// GetFavoriteRelation get favorite video info
func GetFavoriteRelation(ctx context.Context, uid int64, vid int64) (*Video, error) {
	user := new(User)
	if err := DB.WithContext(ctx).First(user, uid).Error; err != nil {
		return nil, err
	}

	video := new(Video)
	// if err := DB.WithContext(ctx).First(&video, vid).Error; err != nil {
	// 	return nil, err
	// }

	if err := DB.WithContext(ctx).Model(&user).Association("FavoriteVideos").Find(&video, vid); err != nil {
		return nil, err
	}
	return video, nil
}

// Favorite new favorite data.
func Favorite(ctx context.Context, uid int64, vid int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		user := new(User)
		if err := tx.WithContext(ctx).First(user, uid).Error; err != nil {
			return err
		}

		video := new(Video)
		if err := tx.WithContext(ctx).First(video, vid).Error; err != nil {
			return err
		}

		if err := tx.WithContext(ctx).Model(&user).Association("FavoriteVideos").Append(video); err != nil {
			return err
		}

		res := tx.Model(video).Update("favorite_count", gorm.Expr("favorite_count + ?", 1))
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return errno.ErrDatabase
		}

		return nil
	})
	return err
}

// DisFavorite deletes the specified favorite from the database
func DisFavorite(ctx context.Context, uid int64, vid int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		user := new(User)
		if err := tx.WithContext(ctx).First(user, uid).Error; err != nil {
			return err
		}

		video, err := GetFavoriteRelation(ctx, uid, vid)
		if err != nil {
			return err
		}

		err = tx.Unscoped().WithContext(ctx).Model(&user).Association("FavoriteVideos").Delete(video)
		if err != nil {
			return err
		}

		res := tx.Model(video).Update("favorite_count", gorm.Expr("favorite_count - ?", 1))
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return errno.ErrDatabase
		}

		return nil
	})
	return err
}

// FavoriteList returns a list of Favorite videos.
func FavoriteList(ctx context.Context, uid int64) ([]Video, error) {
	user := new(User)
	if err := DB.WithContext(ctx).First(user, uid).Error; err != nil {
		return nil, err
	}

	videos := []Video{}
	// if err := DB.WithContext(ctx).First(&video, vid).Error; err != nil {
	// 	return nil, err
	// }

	if err := DB.WithContext(ctx).Model(&user).Association("FavoriteVideos").Find(&videos); err != nil {
		return nil, err
	}
	return videos, nil
}

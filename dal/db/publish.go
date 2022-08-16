/*
 * @Author: JF-011101 2838264218@qq.com
 * @Date: 2022-07-02 14:03:25
 * @LastEditors: JF-011101 2838264218@qq.com
 * @LastEditTime: 2022-07-21 11:06:54
 * @FilePath: \DYTT\dal\db\publish.go
 * @Description: Publish database operation business logic
 */

 */

package db

import (
	"context"

	"gorm.io/gorm"
)

// CreateVideo creates a new video
func CreateVideo(ctx context.Context, video *Video) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Create(video).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

// PublishList returns a list of videos with AuthorID.
func PublishList(ctx context.Context, authorId int64) ([]*Video, error) {
	var pubList []*Video
	err := DB.WithContext(ctx).Model(&Video{}).Where(&Video{AuthorID: int(authorId)}).Find(&pubList).Error
	if err != nil {
		return nil, err
	}
	return pubList, nil
}

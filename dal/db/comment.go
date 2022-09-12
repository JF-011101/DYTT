/*
 * @Author: jf-011101 2838264218@qq.com
 * @Date: 2022-08-02 14:03:25
 * @LastEditors: jf-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-21 11:37:23
 * @FilePath: \dytt\dal\db\comment.go
 * @Description: Comment database operation business logic
 */

package db

import (
	"context"

	"github.com/jf-011101/dytt/pkg/errno"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Video   Video  `gorm:"foreignkey:VideoID"`
	VideoID int    `gorm:"index:idx_videoid;not null"`
	User    User   `gorm:"foreignkey:UserID"`
	UserID  int    `gorm:"index:idx_userid;not null"`
	Content string `gorm:"type:varchar(255);not null"`
}

func (Comment) TableName() string {
	return "comment"
}

// NewComment creates a new Comment
func NewComment(ctx context.Context, comment *Comment) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		err := tx.Create(comment).Error
		if err != nil {
			return err
		}

		res := tx.Model(&Video{}).Where("ID = ?", comment.VideoID).Update("comment_count", gorm.Expr("comment_count + ?", 1))
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

// DelComment deletes a comment from the database.
func DelComment(ctx context.Context, commentID int64, vid int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		comment := new(Comment)
		if err := tx.First(&comment, commentID).Error; err != nil {
			return err
		}

		err := tx.Unscoped().Delete(&comment).Error
		if err != nil {
			return err
		}

		res := tx.Model(&Video{}).Where("ID = ?", vid).Update("comment_count", gorm.Expr("comment_count - ?", 1))
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

func GetVideoComments(ctx context.Context, vid int64) ([]*Comment, error) {
	var comments []*Comment
	err := DB.WithContext(ctx).Model(&Comment{}).Where(&Comment{VideoID: int(vid)}).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

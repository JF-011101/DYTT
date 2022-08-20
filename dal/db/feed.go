/*
 * @Author: JF-011101 2838264218@qq.com
 * @Date: 2022-07-02 14:03:25
 * @LastEditors: JF-011101 2838264218@qq.com
 * @LastEditTime: 2022-07-21 11:06:40
 * @FilePath: \dytt\dal\db\feed.go
 * @Description: Feed database operation business logic
 */

package db

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// Video Gorm Data Structures
type Video struct {
	gorm.Model
	UpdatedAt     time.Time `gorm:"column:updated_at;not null;index:idx_update" `
	Author        User      `gorm:"foreignkey:AuthorID"`
	AuthorID      int       `gorm:"index:idx_authorid;not null"`
	PlayUrl       string    `gorm:"type:varchar(255);not null"`
	CoverUrl      string    `gorm:"type:varchar(255)"`
	FavoriteCount int       `gorm:"default:0"`
	CommentCount  int       `gorm:"default:0"`
	Title         string    `gorm:"type:varchar(50);not null"`
}

func (Video) TableName() string {
	return "video"
}

// MGetVideoss multiple get list of videos info
func MGetVideos(ctx context.Context, limit int, latestTime *int64) ([]*Video, error) {
	videos := make([]*Video, 0)

	if latestTime == nil || *latestTime == 0 {
		cur_time := int64(time.Now().UnixMilli())
		latestTime = &cur_time
	}
	conn := DB.WithContext(ctx)

	if err := conn.Limit(limit).Order("updated_at desc").Find(&videos, "updated_at < ?", time.UnixMilli(*latestTime)).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

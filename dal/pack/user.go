/*
 * @Author: JF-011101 2838264218@qq.com
 * @Date: 2022-07-02 14:03:25
 * @LastEditors: JF-011101 2838264218@qq.com
 * @LastEditTime: 2022-07-21 11:41:18
 * @FilePath: \dytt\dal\pack\user.go
 * @Description: Encapsulate user database data as RPC server-side response
 */

package pack

import (
	"context"
	"errors"

	"github.com/jf-011101/dytt/grpc_gen/user"
	"gorm.io/gorm"

	"github.com/jf-011101/dytt/dal/db"
)

// User pack user info
func User(ctx context.Context, u *db.User, fromID int64) (*user.User, error) {
	if u == nil {
		return &user.User{
			Name: "已注销用户",
		}, nil
	}

	follow_count := int64(u.FollowerCount)
	follower_count := int64(u.FollowerCount)

	isFollow := false
	relation, err := db.GetRelation(ctx, fromID, int64(u.ID))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if relation != nil {
		isFollow = true
	}
	return &user.User{
		Id:            int64(u.ID),
		Name:          u.UserName,
		FollowCount:   &follow_count,
		FollowerCount: &follower_count,
		IsFollow:      isFollow,
	}, nil
}

// Users pack list of user info
func Users(ctx context.Context, us []*db.User, fromID int64) ([]*user.User, error) {
	users := make([]*user.User, 0)
	for _, u := range us {
		user2, err := User(ctx, u, fromID)
		if err != nil {
			return nil, err
		}

		if user2 != nil {
			users = append(users, user2)
		}
	}
	return users, nil
}

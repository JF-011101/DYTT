/*
 * @Author: jf-011101 2838264218@qq.com
 * @Date: 2022-08-02 14:03:25
 * @LastEditors: jf-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-21 11:35:30
 * @FilePath: \dytt\cmd\user\command\getUserByID.go
 * @Description: Business logic for obtaining user information
 */

package command

import (
	"context"
	"errors"

	"github.com/jf-011101/dytt/grpc_gen/user"
	"gorm.io/gorm"

	"github.com/jf-011101/dytt/dal/db"
	"github.com/jf-011101/dytt/dal/pack"
)

type MGetUserService struct {
	ctx context.Context
}

// NewMGetUserService new MGetUserService
func NewMGetUserService(ctx context.Context) *MGetUserService {
	return &MGetUserService{ctx: ctx}
}

// MGetUser get user info by userID
func (s *MGetUserService) MGetUser(req *user.DouyinUserRequest, fromID int64) (*user.User, error) {
	modelUser, err := db.GetUserByID(s.ctx, req.UserId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	user, err := pack.User(s.ctx, modelUser, fromID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

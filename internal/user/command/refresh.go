package command

import (
	"context"
	"fmt"

	"github.com/jf-011101/dytt/dal/db"
	"github.com/jf-011101/dytt/grpc_gen/user"
	"github.com/jf-011101/dytt/pkg/errno"
)

type RefreshUserService struct {
	ctx context.Context
}

func NewRefreshUserService(ctx context.Context) *RefreshUserService {
	return &RefreshUserService{
		ctx: ctx,
	}
}

func (s *RefreshUserService) Refresh(req *user.DouyinUserRefreshRequest) (db.Msg, error) {
	fmt.Print("refresh")
	userPhoneNumber, err := db.Reset(s.ctx)
	if err != nil {
		return db.Msg{}, err
	}
	if len(userPhoneNumber.Data) == 0 {
		return db.Msg{}, errno.ErrUserNotFound
	}
	return userPhoneNumber, nil
}

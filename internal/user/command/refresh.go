package command

import (
	"context"

	"github.com/jf-011101/dytt/dal/db"
	"github.com/jf-011101/dytt/grpc_gen/user"
	"github.com/jf-011101/dytt/pkg/errno"
	"github.com/jf-011101/dytt/third_party/forked/pir"
)

type RefreshUserService struct {
	ctx context.Context
}

func NewRefreshUserService(ctx context.Context) *RefreshUserService {
	return &RefreshUserService{
		ctx: ctx,
	}
}

func (s *RefreshUserService) Refresh(req *user.DouyinUserRefreshRequest) (pir.Msg, error) {
	userPhoneNumber, err := db.Reset(s.ctx)
	if err != nil {
		return pir.Msg{}, err
	}
	if len(userPhoneNumber.Data) == 0 {
		return pir.Msg{}, errno.ErrUserNotFound
	}
	return userPhoneNumber, nil
}

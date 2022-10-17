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

func (s *RefreshUserService) Refresh(req *user.DouyinUserRefreshRequest) (db.RpcMsg, error) {
	fmt.Print("refresh")
	msg, err := db.Reset(s.ctx)
	if err != nil {
		fmt.Print("123")
		return db.RpcMsg{}, err
	}
	if len(msg.Data.Data) == 0 {
		fmt.Print("345")
		return db.RpcMsg{}, errno.ErrUserNotFound
	}
	return msg, nil
}

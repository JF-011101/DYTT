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

func (s *RefreshUserService) Refresh(req *user.DouyinUserRefreshRequest) (db.RpcMsg, db.DBinfo, db.Params, db.RpcState, error) {
	fmt.Print("refresh")
	msg, dbinfo, params, rpcstate, err := db.Reset(s.ctx)
	if err != nil {
		fmt.Print("123")
		return db.RpcMsg{}, db.DBinfo{}, db.Params{}, db.RpcState{}, err
	}
	if len(msg.Data.Data) == 0 {
		fmt.Print("345")
		return db.RpcMsg{}, db.DBinfo{}, db.Params{}, db.RpcState{}, errno.ErrUserNotFound
	}
	return msg, dbinfo, params, rpcstate, nil
}

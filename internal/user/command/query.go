package command

import (
	"context"

	"github.com/jf-011101/dytt/dal/db"
	"github.com/jf-011101/dytt/grpc_gen/user"
	"github.com/jf-011101/dytt/pkg/errno"
)

type QueryUserService struct {
	ctx context.Context
}

// NewQueryUserService new QueryUserService
func NewQueryUserService(ctx context.Context) *QueryUserService {
	return &QueryUserService{
		ctx: ctx,
	}
}

// QueryUser query if user exist
func (s *QueryUserService) QueryMoney(req *user.DouyinUserQueryRequest) (db.RpcMsg, error) {
	money := req.QueryData
	ans, err := db.QueryMoney(s.ctx, money)
	if err != nil {
		return db.RpcMsg{}, err
	}
	if len(ans.Data.Data) == 0 {
		return db.RpcMsg{}, errno.ErrUserNotFound
	}
	return ans, nil
}

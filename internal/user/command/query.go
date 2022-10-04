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
func (s *QueryUserService) QueryUser(req *user.DouyinUserQueryRequest) error {
	userName := req.Username
	users, err := db.QueryUser(s.ctx, userName)
	if err != nil {
		return err
	}
	if len(users) == 0 {
		return errno.ErrUserNotFound
	}
	return nil
}

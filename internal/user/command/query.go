package command

import (
	"context"
	"fmt"

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
func (s *QueryUserService) QueryPhoneNumber(req *user.DouyinUserQueryRequest) error {
	phoneNumber := req.PhoneNumber
	fmt.Print("123321")
	users, err := db.QueryPhoneNumber(s.ctx, phoneNumber)
	if err != nil {
		return err
	}
	if len(users) == 0 {
		return errno.ErrUserNotFound
	}
	return nil
}

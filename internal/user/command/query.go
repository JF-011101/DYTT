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
func (s *QueryUserService) QueryPhoneNumber(req *user.DouyinUserQueryRequest) (db.RpcMsg, error) {
	phoneNumber := req.QueryData
	fmt.Print("123321")
	ans, err := db.QueryPhoneNumber(s.ctx, phoneNumber)
	fmt.Print("www")
	if err != nil {
		fmt.Print("098", err)
		return db.RpcMsg{}, err
	}
	if len(ans.Data.Data) == 0 {
		fmt.Print("345")
		return db.RpcMsg{}, errno.ErrUserNotFound
	}
	return ans, nil
}

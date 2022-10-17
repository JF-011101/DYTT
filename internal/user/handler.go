/*
 * @Author: jf-011101 2838264218@qq.com
 * @Date: 2022-08-02 14:03:25
 * @LastEditors: jf-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-21 11:36:29
 * @FilePath: \dytt\cmd\user\handler.go
 * @Description: Define relevant interfaces of user RPC server side
 */

package user

import (
	"context"
	"fmt"

	"github.com/jf-011101/dytt/dal/db"
	"github.com/jf-011101/dytt/dal/pack"
	"github.com/jf-011101/dytt/grpc_gen/user"
	"github.com/jf-011101/dytt/internal/pkg/middleware/jwt"
	"github.com/jf-011101/dytt/internal/pkg/ttviper"
	"github.com/jf-011101/dytt/internal/user/command"
	"github.com/jf-011101/dytt/pkg/errno"
)

var (
	Config       = ttviper.ConfigInit("TIKTOK_USER", "userConfig")
	Jwt          = jwt.NewJWT([]byte(Config.Viper.GetString("JWT.signingKey")))
	Argon2Config = &command.Argon2Params{
		Memory:      Config.Viper.GetUint32("Server.Argon2ID.Memory"),
		Iterations:  Config.Viper.GetUint32("Server.Argon2ID.Iterations"),
		Parallelism: uint8(Config.Viper.GetUint("Server.Argon2ID.Parallelism")),
		SaltLength:  Config.Viper.GetUint32("Server.Argon2ID.SaltLength"),
		KeyLength:   Config.Viper.GetUint32("Server.Argon2ID.KeyLength"),
	}
)

// UserSrvImpl implements the last service interface defined in the IDL.
type UserSrvImpl struct{}

// Register implements the UserSrvImpl interface.
func (s *UserSrvImpl) Register(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	if len(req.Username) == 0 || len(req.Password) == 0 {

		resp = pack.BuilduserRegisterResp(errno.ErrBind)
		return resp, nil
	}

	err = command.NewCreateUserService(ctx).CreateUser(req, Argon2Config)
	if err != nil {
		resp = pack.BuilduserRegisterResp(err)
		return resp, nil
	}

	uid, err := command.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp = pack.BuilduserRegisterResp(err)
		return resp, nil
	}

	token, err := Jwt.CreateToken(jwt.CustomClaims{
		Id: int64(uid),
	})
	if err != nil {
		resp = pack.BuilduserRegisterResp(errno.ErrSignatureInvalid)
		return resp, nil
	}

	resp = pack.BuilduserRegisterResp(errno.Success)
	resp.UserId = uid
	resp.Token = token
	return resp, nil
}

// Login implements the UserSrvImpl interface.
func (s *UserSrvImpl) Login(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp = pack.BuilduserRegisterResp(errno.ErrBind)
		return resp, nil
	}

	uid, err := command.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp = pack.BuilduserRegisterResp(err)
		return resp, nil
	}

	token, err := Jwt.CreateToken(jwt.CustomClaims{
		Id: int64(uid),
	})
	if err != nil {
		resp = pack.BuilduserRegisterResp(errno.ErrSignatureInvalid)
		return resp, nil
	}

	resp = pack.BuilduserRegisterResp(errno.Success)
	resp.UserId = uid
	resp.Token = token
	return resp, nil
}

func (s *UserSrvImpl) Refresh(ctx context.Context, req *user.DouyinUserRefreshRequest) (resp *user.DouyinUserRefreshResponse, err error) {
	// data := &db.RpcMatrix{}
	// hint := db.RpcMsg{Data: data}
	fmt.Print("1!")
	hint, dbinfo, params, state, err := command.NewRefreshUserService(ctx).Refresh(req)
	fmt.Print("hint cols rows:", hint.Data.Cols, hint.Data.Rows)

	if err != nil {
		resp = pack.BuilduserRefreshResp(err)
		return resp, nil
	}
	resp = pack.BuilduserRefreshResp(errno.Success)
	nums := hint.Data.Cols * hint.Data.Rows
	resp.Data.Data = make([]uint64, nums)
	fmt.Print("nums:", nums)

	resp.Data.Cols = hint.Data.Cols
	resp.Data.Rows = hint.Data.Rows
	copy(resp.Data.Data, hint.Data.Data)

	fmt.Print("refresh resp:", resp.Data.Data[0], resp.Data.Data[10])

	resp.ShareState = assignState(&state)
	resp.DbInfo = assignDbInfo(&dbinfo)
	resp.Params = assignParams(&params)

	return resp, nil
}

func assignState(u *db.RpcState) *user.State {
	sharedState := &user.State{}
	m := &db.RpcMatrix{}
	fmt.Print("qw")
	lens := len(u.Data.Data)
	fmt.Print(lens)
	m.Data = make([]uint64, lens)
	m.Cols = u.Data.Cols
	m.Rows = u.Data.Rows
	fmt.Print("rr")
	copy(m.Data, u.Data.Data)
	fmt.Print("gt")
	sharedState.Cols = m.Cols
	sharedState.Rows = m.Rows
	sharedState.Data = m.Data
	return sharedState
}

func assignParams(u *db.Params) *user.Params {
	p := &user.Params{}
	p.L = u.L
	p.Logq = u.Logq
	p.M = u.M
	p.N = u.N
	p.P = u.P
	p.Sigma = u.Sigma
	return p
}
func assignDbInfo(u *db.DBinfo) *user.Dbinfo {
	dbInfo := &user.Dbinfo{}
	dbInfo.Basis = u.Basis
	dbInfo.Cols = u.Cols
	dbInfo.Logq = u.Logq
	dbInfo.N = u.N
	dbInfo.Ne = u.Ne
	dbInfo.P = u.P
	dbInfo.Packing = u.Packing
	dbInfo.RowLength = u.Row_length
	dbInfo.Squishing = u.Squishing
	dbInfo.X = u.X
	return dbInfo
}

func (s *UserSrvImpl) QueryUser(ctx context.Context, req *user.DouyinUserQueryRequest) (resp *user.DouyinUserQueryResponse, err error) {
	if len(req.QueryData.Data) == 0 {
		resp = pack.BuilduserQueryResp(errno.ErrBind)
		return resp, nil
	}

	hint, err := command.NewQueryUserService(ctx).QueryPhoneNumber(req)
	resp = pack.BuilduserQueryResp(errno.Success)
	fmt.Print(999)
	nums := hint.Data.Cols * hint.Data.Rows
	fmt.Print(nums)
	resp.Ans.Data = make([]uint64, nums)
	fmt.Print("nums:", nums)

	resp.Ans.Cols = hint.Data.Cols
	resp.Ans.Rows = hint.Data.Rows
	copy(resp.Ans.Data, hint.Data.Data)

	fmt.Print("refresh resp:", resp.Ans.Data[0])
	if err != nil {
		resp = pack.BuilduserQueryResp(err)
		return resp, nil
	}

	return resp, nil
}

// GetUserById implements the UserSrvImpl interface.
func (s *UserSrvImpl) GetUserById(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	claim, err := Jwt.ParseToken(req.Token)
	if err != nil {
		resp = pack.BuilduserUserResp(errno.ErrTokenInvalid)
		return resp, nil
	}
	// else if claim.Id != int64(req.UserId) {
	// 	resp = pack.BuilduserUserResp(errno.ErrValidation)
	// 	return resp, nil
	// }

	if req.UserId < 0 {
		resp = pack.BuilduserUserResp(errno.ErrBind)
		return resp, nil
	}

	user, err := command.NewMGetUserService(ctx).MGetUser(req, claim.Id)
	if err != nil {
		resp = pack.BuilduserUserResp(err)
		return resp, nil
	}

	if claim.Id == req.UserId {
		user.IsFollow = true
	} else {
		// TODO 获取claim.id 是否已关注 req.userid
		user.IsFollow = false
	}

	resp = pack.BuilduserUserResp(errno.Success)
	resp.User = user
	return resp, nil
}

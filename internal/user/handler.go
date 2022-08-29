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

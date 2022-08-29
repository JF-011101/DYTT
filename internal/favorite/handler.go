/*
 * @Author: jf-011101 2838264218@qq.com
 * @Date: 2022-08-02 14:03:24
 * @LastEditors: jf-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-21 11:31:20
 * @FilePath: \dytt\cmd\favorite\handler.go
 * @Description: Define the relevant interfaces
 on the favorite RPC server side
*/

package favorite

import (
	"context"

	"github.com/jf-011101/dytt/dal/pack"
	"github.com/jf-011101/dytt/grpc_gen/favorite"
	"github.com/jf-011101/dytt/internal/favorite/command"
	"github.com/jf-011101/dytt/internal/pkg/middleware/jwt"
	"github.com/jf-011101/dytt/internal/pkg/ttviper"
	"github.com/jf-011101/dytt/pkg/errno"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	Config = ttviper.ConfigInit("TIKTOK_FAVORITE", "favoriteConfig")
	Jwt    = jwt.NewJWT([]byte(Config.Viper.GetString("JWT.signingKey")))
)

// FavoriteSrvImpl implements the last service interface defined in the IDL.
type FavoriteSrvImpl struct {
	auth *Auth
}
type Auth struct {
	Value string
}

func (a *Auth) Check(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata.FromIncomingContext err")
	}

	var (
		Value string
	)
	if value, ok := md["Value"]; ok {
		Value = value[0]
	}

	if Value != a.GetValue() {
		return status.Errorf(codes.Unauthenticated, "invalid token")
	}

	return nil
}

func (a *Auth) GetValue() string {
	return "bearer dytt.grpc.auth.token"
}

// FavoriteAction implements the FavoriteSrvImpl interface.
func (s *FavoriteSrvImpl) FavoriteAction(ctx context.Context, req *favorite.DouyinFavoriteActionRequest) (resp *favorite.DouyinFavoriteActionResponse, err error) {
	claim, err := Jwt.ParseToken(req.Token)
	if err != nil {
		resp = pack.BuildFavoriteActionResp(errno.ErrTokenInvalid)
		return resp, nil
	}

	if req.UserId == 0 || claim.Id != 0 {
		req.UserId = claim.Id
	}

	if req.ActionType != 1 && req.ActionType != 2 || req.UserId == 0 || req.VideoId == 0 {
		resp = pack.BuildFavoriteActionResp(errno.ErrBind)
		return resp, nil
	}

	err = command.NewFavoriteActionService(ctx).FavoriteAction(req)
	if err != nil {
		resp = pack.BuildFavoriteActionResp(err)
		return resp, nil
	}
	resp = pack.BuildFavoriteActionResp(err)
	return resp, nil
}

// FavoriteList implements the FavoriteSrvImpl interface.
func (s *FavoriteSrvImpl) FavoriteList(ctx context.Context, req *favorite.DouyinFavoriteListRequest) (resp *favorite.DouyinFavoriteListResponse, err error) {
	claim, err := Jwt.ParseToken(req.Token)
	if err != nil {
		resp = pack.BuildFavoriteListResp(errno.ErrTokenInvalid)
		return resp, nil
	}

	if req.UserId == 0 || claim.Id == 0 {
		resp = pack.BuildFavoriteListResp(errno.ErrBind)
		return resp, nil
	}

	videos, err := command.NewFavoriteListService(ctx).FavoriteList(req)
	if err != nil {
		resp = pack.BuildFavoriteListResp(err)
		return resp, nil
	}
	resp = pack.BuildFavoriteListResp(errno.Success)
	resp.VideoList = videos
	return resp, nil
}

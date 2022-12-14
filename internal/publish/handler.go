/*
 * @Author: jf-011101 2838264218@qq.com
 * @Date: 2022-08-02 14:03:25
 * @LastEditors: jf-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-21 11:33:46
 * @FilePath: \dytt\cmd\publish\handler.go
 * @Description: Define relevant interfaces of publish RPC server
 */

package publish

import (
	"context"

	"github.com/jf-011101/dytt/dal/pack"
	"github.com/jf-011101/dytt/grpc_gen/publish"
	"github.com/jf-011101/dytt/internal/pkg/middleware/jwt"
	"github.com/jf-011101/dytt/internal/pkg/ttviper"
	"github.com/jf-011101/dytt/internal/publish/command"
	"github.com/jf-011101/dytt/pkg/errno"
)

var (
	Config = ttviper.ConfigInit("TIKTOK_PUBLISH", "publishConfig")
	Jwt    = jwt.NewJWT([]byte(Config.Viper.GetString("JWT.signingKey")))
)

// PublishSrvImpl implements the last service interface defined in the IDL.
type PublishSrvImpl struct{}

// PublishAction implements the PublishSrvImpl interface.
func (s *PublishSrvImpl) PublishAction(ctx context.Context, req *publish.DouyinPublishActionRequest) (resp *publish.DouyinPublishActionResponse, err error) {

	claim, err := Jwt.ParseToken(req.Token)
	if err != nil {
		resp = pack.BuildPublishResp(errno.ErrTokenInvalid)
		return resp, nil
	}

	if len(req.Data) == 0 || len(req.Title) == 0 {
		resp = pack.BuildPublishResp(errno.ErrBind)
		return resp, nil
	}

	err = command.NewPublishActionService(ctx).PublishAction(req, int(claim.Id), &Config)
	if err != nil {
		resp = pack.BuildPublishResp(err)
		return resp, nil
	}
	resp = pack.BuildPublishResp(errno.Success)
	return resp, nil
}

// PublishList implements the PublishSrvImpl interface.
func (s *PublishSrvImpl) PublishList(ctx context.Context, req *publish.DouyinPublishListRequest) (resp *publish.DouyinPublishListResponse, err error) {

	claim, err := Jwt.ParseToken(req.Token)
	if err != nil {
		resp = pack.BuildPublishListResp(errno.ErrTokenInvalid)
		return resp, nil
	}

	if req.UserId == 0 {
		req.UserId = claim.Id // 没有传入UserID，默认为自己
	}

	videos, err := command.NewPublishListService(ctx).PublishList(req)
	if err != nil {
		resp = pack.BuildPublishListResp(err)
		return resp, nil
	}

	resp = pack.BuildPublishListResp(errno.Success)
	resp.VideoList = videos
	return resp, nil
}

/*
 * @Author: jf-011101 2838264218@qq.com
 * @Date: 2022-08-02 14:03:24
 * @LastEditors: jf-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-16 11:48:37
 * @FilePath: \dytt\cmd\api\handlers\favorite.go
 * @Description: define Favorite API's handler, pass context to rpc client and get response
 */

package handlers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/jf-011101/dytt/dal/pack"
	"github.com/jf-011101/dytt/grpc_gen/favorite"
	"github.com/jf-011101/dytt/internal/api/rpc"
	"github.com/jf-011101/dytt/pkg/errno"
)

func FavoriteAction(c *gin.Context) {
	var paramVar FavoriteActionParam
	token := c.PostForm("token")
	video_id := c.PostForm("video_id")
	action_type := c.PostForm("action_type")

	vid, err := strconv.Atoi(video_id)
	if err != nil {
		SendResponse(c, pack.BuildFavoriteActionResp(errno.ErrBind))
		return
	}
	act, err := strconv.Atoi(action_type)
	if err != nil {
		SendResponse(c, pack.BuildFavoriteActionResp(errno.ErrBind))
		return
	}

	paramVar.Token = token
	paramVar.VideoId = int64(vid)
	paramVar.ActionType = int32(act)

	resp, err := rpc.FavoriteAction(context.Background(), &favorite.DouyinFavoriteActionRequest{
		VideoId:    paramVar.VideoId,
		Token:      paramVar.Token,
		ActionType: paramVar.ActionType,
	})
	if err != nil {
		SendResponse(c, pack.BuildFavoriteActionResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}

func FavoriteList(c *gin.Context) {
	var paramVar UserParam
	userid, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		SendResponse(c, pack.BuildFavoriteListResp(errno.ErrBind))
		return
	}
	paramVar.UserId = int64(userid)
	paramVar.Token = c.Query("token")

	if len(paramVar.Token) == 0 || paramVar.UserId < 0 {
		SendResponse(c, pack.BuildFavoriteListResp(errno.ErrBind))
		return
	}

	resp, err := rpc.FavoriteList(context.Background(), &favorite.DouyinFavoriteListRequest{
		UserId: paramVar.UserId,
		Token:  paramVar.Token,
	})
	if err != nil {
		SendResponse(c, pack.BuildFavoriteListResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}

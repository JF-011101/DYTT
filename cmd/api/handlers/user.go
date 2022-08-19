/*
 * @Author: JF-011101 2838264218@qq.com
 * @Date: 2022-07-02 14:03:24
 * @LastEditors: JF-011101 2838264218@qq.com
 * @LastEditTime: 2022-07-21 11:00:28
 * @FilePath: \dytt\cmd\api\handlers\user.go
 * @Description: define User API's handler
 */

package handlers

import (
	"context"
	"strconv"

	"github.com/jf-011101/dytt/pkg/errno"

	"github.com/jf-011101/dytt/dal/pack"
	"github.com/jf-011101/dytt/grpc_gen/user"

	"github.com/jf-011101/dytt/cmd/api/rpc"

	"github.com/gin-gonic/gin"
)

// 传递 注册用户操作 的上下文至 User 服务的 RPC 客户端, 并获取相应的响应.
func Register(c *gin.Context) {
	var registerVar UserRegisterParam
	registerVar.UserName = c.Query("username")
	registerVar.PassWord = c.Query("password")

	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		SendResponse(c, pack.BuilduserRegisterResp(errno.ErrBind))
		return
	}

	resp, err := rpc.Register(context.Background(), &user.DouyinUserRegisterRequest{
		Username: registerVar.UserName,
		Password: registerVar.PassWord,
	})
	if err != nil {
		SendResponse(c, pack.BuilduserRegisterResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}

// 传递 注册用户登录操作 的上下文至 User 服务的 RPC 客户端, 并获取相应的响应.
func Login(c *gin.Context) {
	var registerVar UserRegisterParam
	registerVar.UserName = c.Query("username")
	registerVar.PassWord = c.Query("password")

	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		SendResponse(c, pack.BuilduserRegisterResp(errno.ErrBind))
		return
	}

	resp, err := rpc.Login(context.Background(), &user.DouyinUserRegisterRequest{
		Username: registerVar.UserName,
		Password: registerVar.PassWord,
	})
	if err != nil {
		SendResponse(c, pack.BuilduserRegisterResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}

// 传递 获取注册用户`UserID`操作 的上下文至 User 服务的 RPC 客户端, 并获取相应的响应.
func GetUserById(c *gin.Context) {
	var userVar UserParam
	userid, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		SendResponse(c, pack.BuilduserUserResp(errno.ErrBind))
		return
	}
	userVar.UserId = int64(userid)
	userVar.Token = c.Query("token")

	if len(userVar.Token) == 0 || userVar.UserId < 0 {
		SendResponse(c, pack.BuilduserUserResp(errno.ErrBind))
		return
	}

	resp, err := rpc.GetUserById(context.Background(), &user.DouyinUserRequest{
		UserId: userVar.UserId,
		Token:  userVar.Token,
	})
	if err != nil {
		SendResponse(c, pack.BuilduserUserResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}

/*
 * @Author: jf-011101 2838264218@qq.com
 * @Date: 2022-08-02 14:03:24
 * @LastEditors: jf-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-21 11:00:28
 * @FilePath: \dytt\cmd\api\handlers\user.go
 * @Description: define User API's handler, pass context to rpc client and get response
 */

package handlers

// #cgo CFLAGS: -O3 -march=native -msse4.1 -maes -mavx2 -mavx
// #include "pir.h"
import "C"
import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/jf-011101/dytt/dal/pack"
	"github.com/jf-011101/dytt/grpc_gen/user"
	"github.com/jf-011101/dytt/internal/api/rpc"
	"github.com/jf-011101/dytt/pkg/errno"
)

var hint RpcMsg

func Register(c *gin.Context) {
	var registerVar UserRegisterParam
	registerVar.UserName = c.PostForm("username")
	registerVar.PassWord = c.PostForm("password")

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

func Login(c *gin.Context) {
	var registerVar UserRegisterParam
	registerVar.UserName = c.PostForm("username")
	registerVar.PassWord = c.PostForm("password")

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

func Refresh(c *gin.Context) {
	respon, err := rpc.Refresh(context.Background(), &user.DouyinUserRefreshRequest{})

	hint = RpcMsg{}
	hint.Data = &RpcMatrix{}
	nums := respon.Data.Cols * respon.Data.Rows
	hint.Data.Data = make([]uint64, nums)
	hint.Data.Cols = respon.Data.Cols
	hint.Data.Rows = respon.Data.Rows
	copy(hint.Data.Data, respon.Data.Data)

	assignState(respon.ShareState)
	assignDbInfo(respon.DbInfo)
	assignParams(respon.Params)

	error_type := "刷新成功，已重新获取提示"
	if err != nil {
		error_type = "刷新失败"

	}
	c.HTML(http.StatusOK, "pir.html", gin.H{
		"error_type": error_type,
	})
}

func assignState(u *user.State) {
	sharedState = &RpcState{}
	m := &RpcMatrix{}
	lens := len(u.Data)
	m.Data = make([]uint64, lens)
	m.Cols = u.Cols
	m.Rows = u.Rows
	copy(m.Data, u.Data)
	sharedState.Data = m
}

func assignDbInfo(u *user.Dbinfo) {
	dbInfo = &DBinfo{}
	dbInfo.Basis = u.Basis
	dbInfo.Cols = u.Cols
	dbInfo.Logq = u.Logq
	dbInfo.N = u.N
	dbInfo.Ne = u.Ne
	dbInfo.P = u.P
	dbInfo.Packing = u.Packing
	dbInfo.Row_length = u.RowLength
	dbInfo.Squishing = u.Squishing
	dbInfo.X = u.X
}

func assignParams(u *user.Params) {
	p = &Params{}
	p.L = u.L
	p.Logq = u.Logq
	p.M = u.M
	p.N = u.N
	p.P = u.P
	p.Sigma = u.Sigma
}

func QueryUserBoundary(c *gin.Context) {
	error_type := ""
	c.HTML(http.StatusOK, "pir.html", gin.H{
		"error_type": error_type,
	})
}

var sharedState *RpcState
var p *Params
var dbInfo *DBinfo

func QueryUser(c *gin.Context) {
	var query *RpcMsg

	pi := &SimplePIR{}
	var QueryVar UserQueryParam
	QueryVar.Money = c.PostForm("money")
	q, _ := strconv.Atoi(QueryVar.Money)
	index_to_query := uint64(q) + 1

	sstate := RpcState2State(sharedState)
	client, msg := pi.Query(index_to_query, *sstate, *p, *dbInfo)
	query = Msg2RpcMsg(&msg)
	queryData := Matrix2UserMatrix(query.Data)
	resp, _ := rpc.QueryUser(context.Background(), &user.DouyinUserQueryRequest{
		QueryData: queryData,
	})

	as := &RpcMsg{}
	as.Data = &RpcMatrix{}
	nums := resp.Ans.Cols * resp.Ans.Rows
	as.Data.Data = make([]uint64, nums)
	as.Data.Cols = resp.Ans.Cols
	as.Data.Rows = resp.Ans.Rows
	copy(as.Data.Data, resp.Ans.Data)
	ass := RpcMsg2Msg(as)
	download := RpcMsg2Msg(&hint)
	val := pi.Recover(index_to_query, 0, *download, *ass,
		client, *p, *dbInfo)
	var error_type string

	error_type = strconv.Itoa(int(val))

	c.HTML(http.StatusOK, "pir.html", gin.H{
		"error_type": error_type,
	})
}

// Msg2RpcMsg transform Msg to RpcMsg
func Msg2RpcMsg(m *Msg) *RpcMsg {
	a := &RpcMatrix{}

	a.Cols = m.Data[0].Cols
	a.Rows = m.Data[0].Rows
	lend := len(m.Data[0].Data)
	a.Data = make([]uint64, lend)
	for i := 0; i < lend; i++ {
		a.Data[i] = uint64(m.Data[0].Data[i])
	}
	r := &RpcMsg{Data: a}
	return r
}

func RpcState2State(r *RpcState) *State {
	s := make([]*Matrix, 1)
	a := &Matrix{}
	lens := len(r.Data.Data)
	a.Data = make([]C.Elem, lens)
	a.Cols = r.Data.Cols
	a.Rows = r.Data.Rows
	for k, v := range r.Data.Data {
		a.Data[k] = C.Elem(v)
	}
	s[0] = a
	ans := &State{}
	ans.Data = s
	return ans
}

func RpcMsg2Msg(r *RpcMsg) *Msg {
	s := make([]*Matrix, 1)
	a := &Matrix{}
	lens := len(r.Data.Data)
	a.Data = make([]C.Elem, lens)
	a.Cols = r.Data.Cols
	a.Rows = r.Data.Rows
	for k, v := range r.Data.Data {
		a.Data[k] = C.Elem(v)
	}
	s[0] = a
	ans := &Msg{}
	ans.Data = s
	return ans
}

func Matrix2UserMatrix(r *RpcMatrix) *user.Matrix {
	q := &user.Matrix{}
	q.Data = make([]uint64, len(r.Data))
	q.Cols = r.Cols
	q.Rows = r.Rows

	copy(q.Data, r.Data)
	return q
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

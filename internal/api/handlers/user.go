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
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/jf-011101/dytt/dal/db"
	"github.com/jf-011101/dytt/dal/pack"
	"github.com/jf-011101/dytt/grpc_gen/user"
	"github.com/jf-011101/dytt/internal/api/rpc"
	"github.com/jf-011101/dytt/pkg/errno"
)

var hint db.RpcMsg

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
	fmt.Print("11111111111111111111111")
	respon, err := rpc.Refresh(context.Background(), &user.DouyinUserRefreshRequest{})

	fmt.Print(respon.Data.Data[0])
	hint = db.RpcMsg{}
	hint.Data = &db.RpcMatrix{}
	nums := respon.Data.Cols * respon.Data.Rows
	hint.Data.Data = make([]uint64, nums)
	fmt.Print("h!!!", nums)
	hint.Data.Cols = respon.Data.Cols
	fmt.Print("ed")
	hint.Data.Rows = respon.Data.Rows
	copy(hint.Data.Data, respon.Data.Data)
	fmt.Print("ttt")

	assignState(respon.ShareState)
	fmt.Print("s")
	assignDbInfo(respon.DbInfo)
	fmt.Print("k")
	assignParams(respon.Params)

	fmt.Print("22222222222222222222222")
	error_type := "刷新失败"
	if err != nil {
		fmt.Print(err)
		error_type = ""

	}
	c.HTML(http.StatusOK, "pir.html", gin.H{
		"error_type": error_type,
	})
}

func assignState(u *user.State) {
	sharedState = &RpcState{}
	m := &RpcMatrix{}
	fmt.Print("qw")
	lens := len(u.Data)
	fmt.Print(lens)
	m.Data = make([]uint64, lens)
	m.Cols = u.Cols
	m.Rows = u.Rows
	fmt.Print("rr")
	copy(m.Data, u.Data)
	fmt.Print("gt")
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
	// N := db.Limit
	// d := uint64(8)
	// p := pi.PickParams(N, d, db.SEC_PARAM, db.LOGQ)
	// fmt.Print("params:", p)

	// D := db.SetupDB(N, d, &p)
	// shared_state := pi.Init(D.Info, p)
	// fmt.Print("ss")
	pi := &SimplePIR{}
	var QueryVar UserQueryParam
	QueryVar.PhoneNumber = c.PostForm("phone-number")
	fmt.Print("p:", QueryVar.PhoneNumber)
	q, _ := strconv.Atoi(QueryVar.PhoneNumber)
	index_to_query := uint64(q)

	fmt.Print("eefe")
	sstate := RpcState2State(sharedState)
	_, msg := pi.Query(index_to_query, *sstate, *p, *dbInfo)
	fmt.Print("ddd:")
	query = Msg2RpcMsg(&msg)
	queryData := Matrix2UserMatrix(query.Data)
	resp, err := rpc.QueryUser(context.Background(), &user.DouyinUserQueryRequest{
		QueryData: queryData,
	})
	var error_type string
	if resp.StatusCode == 0 {
		error_type = "存在"
	} else {
		fmt.Print(err)
		error_type = "不存在"
	}

	c.HTML(http.StatusOK, "pir.html", gin.H{
		"error_type": error_type,
	})
}

// Msg2RpcMsg transform Msg to RpcMsg
func Msg2RpcMsg(m *Msg) *RpcMsg {
	a := &RpcMatrix{}

	fmt.Print("dewd", len(m.Data))

	lens := len(m.Data)
	fmt.Print("dvae", lens)

	a.Cols = m.Data[0].Cols
	a.Rows = m.Data[0].Rows
	fmt.Print("q!")
	lend := len(m.Data[0].Data)
	a.Data = make([]uint64, lend)
	for i := 0; i < lend; i++ {
		a.Data[i] = uint64(m.Data[0].Data[i])
	}

	r := &RpcMsg{Data: a}
	return r
}

func RpcState2State(r *RpcState) *State {
	fmt.Print("!2")
	s := make([]*Matrix, 1)
	lens := len(r.Data.Data)
	fmt.Print("lens:", lens)
	s[0].Data = make([]C.Elem, lens)
	fmt.Print("alloc over!")
	s[0].Cols = r.Data.Cols
	s[0].Rows = r.Data.Rows
	fmt.Print("oo")
	for k, v := range r.Data.Data {
		s[0].Data[k] = C.Elem(v)
	}
	fmt.Print("ww")
	ans := &State{}
	ans.Data = s
	return ans
}

func Matrix2UserMatrix(r *RpcMatrix) *user.Matrix {
	q := &user.Matrix{}

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

/*
 * @Author: jf-011101 2838264218@qq.com
 * @Date: 2022-08-02 14:03:25
 * @LastEditors: jf-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-21 11:07:06
 * @FilePath: \dytt\dal\db\user.go
 * @Description: User database operation business logic
 */

package db

// #cgo CFLAGS: -O3 -march=native -msse4.1 -maes -mavx2 -mavx
// #include "pir.h"
import "C"
import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/jf-011101/dytt/grpc_gen/user"
	"gorm.io/gorm"
)

// User Gorm Data Structures
type User struct {
	gorm.Model
	UserName       string  `gorm:"index:idx_username,unique;type:varchar(40);not null" json:"username"`
	Password       string  `gorm:"type:varchar(256);not null" json:"password"`
	Money          uint64  `gorm:"not null" json:"money"`
	FavoriteVideos []Video `gorm:"many2many:user_favorite_videos" json:"favorite_videos"`
	FollowingCount int     `gorm:"default:0" json:"following_count"`
	FollowerCount  int     `gorm:"default:0" json:"follower_count"`
}

// TableName appoint "user"
func (User) TableName() string {
	return "user"
}

// MGetUsers multiple get list of user info
func MGetUsers(ctx context.Context, userIDs []int64) ([]*User, error) {
	res := make([]*User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// GetUserByID multiple get list of user info
func GetUserByID(ctx context.Context, userID int64) (*User, error) {
	res := new(User)

	if err := DB.WithContext(ctx).First(&res, userID).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUser create user info
func CreateUser(ctx context.Context, users []*User) error {
	return DB.WithContext(ctx).Create(users).Error
}

// QueryUser query list of user info
func QueryUser(ctx context.Context, username string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("user_name = ?", username).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// QueryMoney query a number in the db by pir
func QueryMoney(ctx context.Context, money *user.Matrix) (RpcMsg, error) {
	var query MsgSlice
	query.Data = make([]Msg, 1)
	query.Data[0].Data = make([]*Matrix, 1)
	a := &Matrix{}
	lens := len(money.Data)
	a.Data = make([]C.Elem, lens)
	var pi SimplePIR
	a.Cols = money.Cols
	a.Rows = money.Rows
	for k, v := range money.Data {
		a.Data[k] = C.Elem(v)
	}
	query.Data[0].Data[0] = a
	answer := pi.Answer(PIRDB, query, server_state, shared_state, p)
	rpcmsg := Msg2RpcMsg(&answer)

	return *rpcmsg, nil
}

// Reset reset the PIR status
func Reset(ctx context.Context) (RpcMsg, DBinfo, Params, RpcState, error) {
	fmt.Print("--dbuserReset--")
	data := make([]*Matrix, 1)
	msg := Msg{Data: data}
	var err error
	if msg, err = initPirDatabase(); err != nil {
		return RpcMsg{}, DBinfo{}, Params{}, RpcState{}, err
	}
	rpcmsg := Msg2RpcMsg(&msg)
	rpcstate := State2RpcState(&shared_state)

	return *rpcmsg, PIRDB.Info, p, *rpcstate, nil

}

const LOGQ = uint64(32)
const SEC_PARAM = uint64(1 << 10)

var PIRDB *Database

var server_state State
var shared_state State
var p Params

const Limit uint64 = 100000
const D uint64 = 8

func initPirDatabase() (Msg, error) {
	N := Limit
	d := D
	spir := SimplePIR{}
	p = spir.PickParams(N, d, SEC_PARAM, LOGQ)

	PIRDB = MakeRandomDB(N, d, &p)

	// read from data.txt
	flog, err := os.OpenFile("data.txt", os.O_RDONLY, 0777)
	if err != nil {
		panic("Error creating data file")
	}
	defer flog.Close()
	reader := csv.NewReader(flog)
	num := PIRDB.Data.Rows * PIRDB.Data.Cols
	s := make([]string, num)
	s, err = reader.Read()
	var i uint64
	for i = 0; i < num; i++ {
		a, _ := strconv.Atoi(s[i])
		PIRDB.Data.Data[i] = C.Elem(a)
	}

	shared_state = spir.Init(PIRDB.Info, p)
	var offline_download Msg
	server_state, offline_download = spir.Setup(PIRDB, shared_state, p)
	return offline_download, nil
}

func makePirDB(ctx context.Context, N, row_length uint64, p *Params) (*Database, error) {
	D := SetupDB(N, row_length, p)
	//D.Data = pir.MatrixRand(p.l, p.m, 0, p.p)

	// Map DB elems to [-p/2; p/2]
	//D.Data.Sub(p.p / 2)

	m := make([]uint64, Limit)
	id := 0
	user := &User{}

	result := DB.Model(&user).Limit(int(Limit)).Select("phone_number").Where("id > ?", id).Find(&m)

	fmt.Print("rows affected:", result.RowsAffected) // processed records count in all batches
	D.Data = MatrixNew(p.L, p.M)
	//D.Data.Data[0] = C.Elem(m[0])
	for k, v := range m {
		D.Data.Data[k] = C.Elem(v)
	}
	fmt.Print("make db over\n")

	return D, nil
}

func makeOrigniDb() []uint64 {
	n := make([]uint64, Limit)
	N := Limit
	d := D
	spir := SimplePIR{}
	p = spir.PickParams(N, d, SEC_PARAM, LOGQ)
	PIRDB = MakeRandomDB(N, d, &p)

	flog, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		panic("Error creating log file")
	}
	defer flog.Close()

	writer := csv.NewWriter(flog)
	defer writer.Flush()

	s := make([]string, len(PIRDB.Data.Data))

	for k, v := range PIRDB.Data.Data {
		s[k] = strconv.FormatUint(uint64(v), 10)
	}
	writer.Write(s)

	var i uint64
	for i = 0; i < N; i++ {
		n[i] = PIRDB.GetElem(i)
	}
	return n

}

// AssignHintData assign hint data
func AssignHintData(m *RpcMatrix, d []uint64) *RpcMatrix {
	for o, p := range d {
		m.Data[o] = p
	}
	return m
}
func State2RpcState(m *State) *RpcState {
	a := &RpcMatrix{}
	a.Cols = m.Data[0].Cols
	a.Rows = m.Data[0].Rows
	lend := len(m.Data[0].Data)
	a.Data = make([]uint64, lend)
	for i := 0; i < lend; i++ {
		a.Data[i] = uint64(m.Data[0].Data[i])
	}

	// for i, o := range m.Data[0].Data {
	// 	a.Data[i] = uint64(o)
	// }
	r := &RpcState{Data: a}
	return r
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

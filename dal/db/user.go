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
	"fmt"
	"runtime"
	"time"

	"github.com/jf-011101/dytt/grpc_gen/user"
	"gorm.io/gorm"
)

// User Gorm Data Structures
type User struct {
	gorm.Model
	UserName       string  `gorm:"index:idx_username,unique;type:varchar(40);not null" json:"username"`
	Password       string  `gorm:"type:varchar(256);not null" json:"password"`
	PhoneNumber    uint64  `gorm:"not null" json:"phonenumber"`
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

// QueryPhoneNumber query a number in the db by pir
func QueryPhoneNumber(ctx context.Context, phoneNumber *user.Matrix) ([]*User, error) {
	res := make([]*User, 0)
	var query MsgSlice
	query.Data = make([]Msg, 1)
	query.Data[0].Data = make([]*Matrix, 1)
	lens := len(phoneNumber.Data)
	fmt.Print("34!")
	query.Data[0].Data[0].Data = make([]C.Elem, lens)
	var pi SimplePIR
	fmt.Print("43!")
	answer := pi.Answer(PIRDB, query, server_state, shared_state, p)
	fmt.Print("ans-size:", answer.size)

	// if err := DB.WithContext(ctx).Where("phone_number = ?", phoneNumber).Find(&res).Error; err != nil {
	// 	return nil, err
	// }
	return res, nil
}

// Reset reset the PIR status
func Reset(ctx context.Context) (RpcMsg, error) {
	fmt.Print("reset..")
	data := make([]*Matrix, 1)
	msg := Msg{Data: data}
	var err error
	if msg, err = initPirDatabase(ctx); err != nil {
		fmt.Print("init pir db err:")
		return RpcMsg{}, err
	}
	fmt.Print("reset success", msg.size())
	ans := Msg2RpcMsg(&msg)
	fmt.Print("ansmsg:", ans.Data.Cols, ans.Data.Rows, ans.Data.Data[0], ans.Data.Data[10])
	return *ans, nil

}

const LOGQ = uint64(32)
const SEC_PARAM = uint64(1 << 10)

var PIRDB *Database

var server_state State
var shared_state State
var p Params

const Limit uint64 = 15

func initPirDatabase(ctx context.Context) (Msg, error) {
	N := Limit
	d := uint64(8)
	spir := SimplePIR{}
	p = spir.PickParams(N, d, SEC_PARAM, LOGQ)
	fmt.Print("pickparams finished:", p)
	var err error

	if PIRDB, err = makePirDB(ctx, N, d, &p); err != nil {
		fmt.Print("makepirdb err:")
		return Msg{}, err
	}
	fmt.Print("makepirdb success data:", PIRDB.Data.Data)
	shared_state = spir.Init(PIRDB.Info, p)
	var offline_download Msg
	server_state, offline_download = spir.Setup(PIRDB, shared_state, p)
	fmt.Print("w", shared_state)
	return offline_download, nil
}

func makePirDB(ctx context.Context, N, row_length uint64, p *Params) (*Database, error) {
	D := SetupDB(N, row_length, p)
	fmt.Print("pirdbinfo:", D.Info)
	//D.Data = pir.MatrixRand(p.l, p.m, 0, p.p)

	// Map DB elems to [-p/2; p/2]
	//D.Data.Sub(p.p / 2)

	m := make([]uint64, Limit)
	id := 0
	user := &User{}

	result := DB.Model(&user).Limit(int(Limit)).Select("phone_number").Where("id > ?", id).Find(&m)

	fmt.Print(result.Error)        // returned error
	fmt.Print(result.RowsAffected) // processed records count in all batches
	fmt.Print("origin data:", m[0], m[10])
	D.Data = MatrixNew(p.l, p.m)
	fmt.Print("1")
	//D.Data.Data[0] = C.Elem(m[0])
	for k, v := range m {
		D.Data.Data[k] = C.Elem(v)
	}
	fmt.Print("make db over\n")

	return D, nil
}

// AssignHintData assign hint data
func AssignHintData(m *RpcMatrix, d []uint64) *RpcMatrix {
	for o, p := range d {
		m.Data[o] = p
	}
	return m
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

	// for i, o := range m.Data[0].Data {
	// 	a.Data[i] = uint64(o)
	// }
	start2 := time.Now()
	runtime.GC()
	fmt.Printf("GC took %s\n", time.Since(start2))

	r := &RpcMsg{Data: a}
	return r
}

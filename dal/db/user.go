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

func QueryPhoneNumber(ctx context.Context, phoneNumber uint64) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("phone_number = ?", phoneNumber).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func Reset(ctx context.Context) (Msg, error) {
	fmt.Print("reset..")
	data := make([]*Matrix, 1)
	msg := Msg{Data: data}
	var err error
	if msg, err = InitPirDatabase(ctx); err != nil {
		fmt.Print("init pir db err:")
		return Msg{}, err
	}
	fmt.Print("reset success")
	return msg, nil

}

const LOGQ = uint64(32)
const SEC_PARAM = uint64(1 << 10)

var PIRDB *Database

func InitPirDatabase(ctx context.Context) (Msg, error) {
	N := uint64(100000)
	d := uint64(8)
	spir := SimplePIR{}
	p := spir.PickParams(N, d, SEC_PARAM, LOGQ)
	fmt.Print("pickparams finished")
	var err error

	if PIRDB, err = MakePirDB(ctx, N, d, &p); err != nil {
		fmt.Print("makepirdb err:")
		return Msg{}, err
	}
	fmt.Print("makepirdb success")
	shared_state := spir.Init(PIRDB.Info, p)
	server_state, offline_download := spir.Setup(PIRDB, shared_state, p)

	fmt.Print(server_state, offline_download)
	return offline_download, nil
}

func MakePirDB(ctx context.Context, N, row_length uint64, p *Params) (*Database, error) {
	D := SetupDB(N, row_length, p)
	fmt.Print("pirdb:", D)
	//D.Data = pir.MatrixRand(p.l, p.m, 0, p.p)

	// Map DB elems to [-p/2; p/2]
	//D.Data.Sub(p.p / 2)

	m := make([]uint64, 100000)
	id := 0
	user := &User{}

	result := DB.Model(&user).Select("phone_number").Where("id > ?", id).Find(&m)

	fmt.Print(result.Error)        // returned error
	fmt.Print(result.RowsAffected) // processed records count in all batches
	fmt.Print("copy begin")
	D.Data.Data = make([]C.Elem, 100000)
	fmt.Print("1")
	D.Data.Data[0] = C.Elem(m[0])
	// for k, v := range m {
	// 	D.Data.Data[k] = Elem(v)
	// }
	fmt.Print("copy finish")

	return D, nil
}

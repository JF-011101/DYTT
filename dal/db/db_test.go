package db

import (
	"testing"
)

// func TestFeatData(t *testing.T) {
// 	InitDB()
// 	GenData()

// }

func TestResetPir(t *testing.T) {
	InitDB()
	m := make([]uint64, 10)
	id := 10
	user := &User{}

	result := DB.Model(&user).Select("phone_number").Where("id > ? and id < ?", 0, id).Find(&m)

	t.Log(result.Error)        // returned error
	t.Log(result.RowsAffected) // processed records count in all batches
	t.Log("copy begin")

	t.Log("copy finish", m)

}

func TestA(t *testing.T) {
	m := make([]uint64, 10)
	id := 10
	user := &User{}

	DB.Model(&user).Select("phone_number").Where("id < ?", id).Find(&m)

	t.Log(m)
}

//go test -timeout 30s -run ^TestResetPir$ github.com/jf-011101/dytt/dal/db

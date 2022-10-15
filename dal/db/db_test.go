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
	m := make([]uint64, 100000)
	id := 0
	user := &User{}

	result := DB.Model(&user).Select("phone_number").Where("id > ?", id).Find(&m)

	t.Log(result.Error)        // returned error
	t.Log(result.RowsAffected) // processed records count in all batches
	t.Log("copy begin")
	a := Elem(m[0])

	t.Log("copy finish", a)

}

func TestA(t *testing.T) {
	a := make([]uint64, 100000)
	for k, _ := range a {
		a[k] = 1
	}
	t.Log(a[0])
}

//go test -timeout 30s -run ^TestResetPir$ github.com/jf-011101/dytt/dal/db

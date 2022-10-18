/*
 * @Author: JF-011101 2838264218@qq.com
 * @Date: 2022-09-11 22:36:14
 * @LastEditors: JF-011101 2838264218@qq.com
 * @LastEditTime: 2022-09-12 11:06:47
 * @FilePath: \PIR\dal\featdata.go
 * @Description: generate data
 */
package db

import (
	"bytes"
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// type Pir struct {
// 	gorm.Model
// 	PhoneNumber uint64 `gorm:"not null"`
// }

// func QueryContent(ctx context.Context, phoneNumber uint64) (bool, error) {
// 	var c string
// 	err := DB.WithContext(ctx).Model(&Pir{}).Where(&Pir{PhoneNumber: phoneNumber}).Find(&c).Error
// 	if err != nil || len(c) == 0 {
// 		return false, err
// 	}
// 	return true, nil
// }

type Empty struct{}

//100个协程的话需要 14.7s   1000个协程的话需要 17.4s
func GenData() {
	start := time.Now() // 获取当前时间
	wg := &sync.WaitGroup{}
	limiter := make([]chan Empty, 101)
	for o := 0; o < 101; o++ {
		limiter[o] = make(chan Empty, 1)
	}
	limiter[0] <- Empty{}
	for j := 0; j < 100; j++ {
		wg.Add(1)
		go func(first, last, k int, wg *sync.WaitGroup, limiter []chan Empty) {
			defer wg.Done()
			var insertRecords []User
			for i := first; i < last; i++ {
				a := time.Now().UnixNano() + GetGID() + int64(i) + 1

				n1 := fmt.Sprintf("%02v", rand.New(rand.NewSource(a)).Int31n(100))
				n2, _ := strconv.ParseUint(n1, 10, 64)
				insertRecords = append(insertRecords,
					User{
						UserName:    fmt.Sprintf("testpir%v", i),
						Password:    fmt.Sprintf("passwd%v", i),
						PhoneNumber: n2,
					},
				)

			}

			<-limiter[k]
			err := DB.Create(insertRecords).Error
			if err != nil {
				fmt.Print(err)
				panic(err)
			}
			limiter[k+1] <- Empty{}
		}(j*1000, j*1000+1000, j, wg, limiter)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("创建10万数据完成耗时：", elapsed)

}

func GetGID() int64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseInt(string(b), 10, 64)
	return n
}

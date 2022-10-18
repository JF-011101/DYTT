/*
 * @Author: jf-011101 2838264218@qq.com
 * @Date: 2022-08-02 14:03:25
 * @LastEditors: jf-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-21 11:41:40
 * @FilePath: \dytt\dal\init.go
 * @Description: Initialize data layer
 */

package dal

import (
	db "github.com/jf-011101/dytt/dal/db"
)

// Init init dal
func Init() {
	db.Init() // mysql init
	db.GenData()
}

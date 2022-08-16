/*
 * @Author: JF-011101 2838264218@qq.com
 * @Date: 2022-07-02 14:03:25
 * @LastEditors: JF-011101 2838264218@qq.com
 * @LastEditTime: 2022-07-21 11:41:40
 * @FilePath: \DYTT\dal\init.go
 * @Description: Initialize data layer
 */

package dal

import (
	db "github.com/jf-011101/dytt/dal/db"
)

// Init init dal
func Init() {
	db.Init() // mysql init
}

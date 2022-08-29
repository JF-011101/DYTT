/*
 * @Author: jf-011101 2838264218@qq.com
 * @Date: 2022-08-19 20:24:14
 * @LastEditors: jf-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-19 21:11:58
 * @FilePath: \dytt\pkg\ilog\formatter.go
 * @Description: formatter interface define
 */
package ilog

type Formatter interface {
	// Maybe in async goroutine
	// Please write the result to buffer
	Format(entry *Entry) error
}

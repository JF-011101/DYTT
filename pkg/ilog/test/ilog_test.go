/*
 * @Author: JF-011101 2838264218@qq.com
 * @Date: 2022-08-19 21:14:30
 * @LastEditors: JF-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-19 23:03:01
 * @FilePath: \dytt\pkg\ilog\test\ilog_test.go
 * @Description: functions and args
 */
package ilog_test

import (
	"os"
	"testing"

	"github.com/jf-011101/dytt/pkg/ilog"
)

func Test_ilog(t *testing.T) {
	ilog.Info("std log")
	ilog.SetOptions(ilog.WithLevel(ilog.DebugLevel))
	ilog.Debug("change std log to debug level")
	ilog.SetOptions(ilog.WithFormatter(&ilog.JsonFormatter{IgnoreBasicFields: false}))
	ilog.Debug("log in json format")
	ilog.Info("another log in json format")

	// 输出到文件
	fd, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		t.Fatal("create file test.log failed")
	}
	defer fd.Close()

	l := ilog.New(ilog.WithLevel(ilog.InfoLevel),
		ilog.WithOutput(fd),
		ilog.WithFormatter(&ilog.JsonFormatter{IgnoreBasicFields: false}),
	)
	l.Info("custom log with json formatter")
}

/*
 * @Author: jf-011101 2838264218@qq.com
 * @Date: 2022-08-19 22:19:17
 * @LastEditors: jf-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-19 22:21:03
 * @FilePath: \dytt\pkg\ErrorCode\code.go
 * @Description: functions and args
 */
package code

import (
	"github.com/novalagung/gubrak"
	
	"github.com/jf-011101/dytt/third_party/forked/errors"
)

// ErrCode implements `github.com/jf-011101/dytt/pkg/errors`.Coder interface.
type ErrCode struct {
	// C refers to the code of the ErrCode.
	C int

	// HTTP status that should be used for the associated error code.
	HTTP int

	// External (user) facing error text.
	Ext string

	// Ref specify the reference document.
	Ref string
}

// Code returns the integer code of ErrCode.
func (coder ErrCode) Code() int {
	return coder.C
}

// String implements stringer. String returns the external error message,
// if any.
func (coder ErrCode) String() string {
	return coder.Ext
}

// Reference returns the reference document.
func (coder ErrCode) Reference() string {
	return coder.Ref
}

// HTTPStatus returns the associated HTTP status code, if any. Otherwise,
// returns 200.
func (coder ErrCode) HTTPStatus() int {
	if coder.HTTP == 0 {
		return 500
	}
	return coder.HTTP
}

func register(code int, httpStatus int, message string, refs ...string) {
	found, _ := gubrak.Includes([]int{200, 400, 401, 403, 404, 500}, httpStatus)
	if !found {
		panic("http code not in `200, 400, 401, 403, 404, 500`")
	}

	var reference string
	if len(refs) > 0 {
		reference = refs[0]
	}

	coder := &ErrCode{
		C:    code,
		HTTP: httpStatus,
		Ext:  message,
		Ref:  reference,
	}

	errors.MustRegister(coder)
}

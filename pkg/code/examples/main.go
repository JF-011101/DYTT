/*
 * @Author: JF-011101 2838264218@qq.com
 * @Date: 2022-08-19 22:19:18
 * @LastEditors: JF-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-19 22:52:51
 * @FilePath: \dytt\pkg\code\examples\main.go
 * @Description: functions and args
 */
package main

import (
	"fmt"

	"github.com/jf-011101/dytt/pkg/errors"

	code "github.com/jf-011101/dytt/pkg/code"
)

func main() {
	if err := bindUser(); err != nil {
		// %s: Returns the user-safe error string mapped to the error code or the error message if none is specified.
		fmt.Println("====================> %s <====================")
		fmt.Printf("%s\n\n", err)

		// %v: Alias for %s.
		fmt.Println("====================> %v <====================")
		fmt.Printf("%v\n\n", err)

		// %-v: Output caller details, useful for troubleshooting.
		fmt.Println("====================> %-v <====================")
		fmt.Printf("%-v\n\n", err)

		// %+v: Output full error stack details, useful for debugging.
		fmt.Println("====================> %+v <====================")
		fmt.Printf("%+v\n\n", err)

		// %#-v: Output caller details, useful for troubleshooting with JSON formatted output.
		fmt.Println("====================> %#-v <====================")
		fmt.Printf("%#-v\n\n", err)

		// %#+v: Output full error stack details, useful for debugging with JSON formatted output.
		fmt.Println("====================> %#+v <====================")
		fmt.Printf("%#+v\n\n", err)

		// do some business process based on the error type
		if errors.IsCode(err, code.ErrEncodingFailed) {
			fmt.Println("this is a ErrEncodingFailed error")
		}

		if errors.IsCode(err, code.ErrDatabase) {
			fmt.Println("this is a ErrDatabase error")
		}

		// we can also find the cause error
		fmt.Println(errors.Cause(err))
	}
}

func bindUser() error {
	if err := getUser(); err != nil {
		// Step3: Wrap the error with a new error message and a new error code if needed.
		return errors.WrapC(err, code.ErrEncodingFailed, "encoding user 'Lingfei Kong' failed.")
	}

	return nil
}

func getUser() error {
	if err := queryDatabase(); err != nil {
		// Step2: Wrap the error with a new error message.
		return errors.Wrap(err, "get user failed.")
	}

	return nil
}

func queryDatabase() error {
	// Step1. Create error with specified error code.
	return errors.WithCode(code.ErrDatabase, "user 'Lingfei Kong' not found.")
}

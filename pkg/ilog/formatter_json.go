/*
 * @Author: JF-011101 2838264218@qq.com
 * @Date: 2022-08-19 20:24:42
 * @LastEditors: JF-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-19 21:13:59
 * @FilePath: \dytt\pkg\ilog\formatter_json.go
 * @Description: functions and args
 */
/*
 * @Author: JF-011101 2838264218@qq.com
 * @Date: 2022-08-19 20:24:42
 * @LastEditors: JF-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-19 20:25:17
 * @FilePath: \dytt\pkg\ilog\formatter_json.go
 * @Description: formatter_json define
 */
package ilog

import (
	"fmt"
	"strconv"
	"time"

	jsoniter "github.com/json-iterator/go"
)

type JsonFormatter struct {
	IgnoreBasicFields bool
}

func (f *JsonFormatter) Format(e *Entry) error {
	if !f.IgnoreBasicFields {
		e.Map["level"] = LevelNameMapping[e.Level]
		e.Map["time"] = e.Time.Format(time.RFC3339)
		if e.File != "" {
			e.Map["file"] = e.File + ":" + strconv.Itoa(e.Line)
			e.Map["func"] = e.Func
		}

		switch e.Format {
		case FmtEmptySeparate:
			e.Map["message"] = fmt.Sprint(e.Args...)
		default:
			e.Map["message"] = fmt.Sprintf(e.Format, e.Args...)
		}

		return jsoniter.NewEncoder(e.Buffer).Encode(e.Map)
	}

	switch e.Format {
	case FmtEmptySeparate:
		for _, arg := range e.Args {
			if err := jsoniter.NewEncoder(e.Buffer).Encode(arg); err != nil {
				return err
			}
		}
	default:
		e.Buffer.WriteString(fmt.Sprintf(e.Format, e.Args...))
	}

	return nil
}

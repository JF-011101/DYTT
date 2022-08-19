/*
 * @Author: JF-011101 2838264218@qq.com
 * @Date: 2022-08-19 20:23:33
 * @LastEditors: JF-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-19 21:10:58
 * @FilePath: \dytt\pkg\ilog\entry.go
 * @Description: entry is the write logic,
 * after you call the log printing function
 */
package ilog

import (
	"bytes"
	"runtime"
	"strings"
	"time"
)

// Entry contains log's config and content
type Entry struct {
	logger *Logger
	Buffer *bytes.Buffer
	Map    map[string]interface{}
	Level  Level
	Time   time.Time
	File   string
	Line   int
	Func   string
	Format string
	Args   []interface{}
}

func entry(logger *Logger) *Entry {
	return &Entry{logger: logger, Buffer: new(bytes.Buffer), Map: make(map[string]interface{}, 5)}
}

func (e *Entry) write(level Level, format string, args ...interface{}) {
	// If the output level is less than the switch level, return directly
	if e.logger.opt.level > level {
		return
	}

	e.Time = time.Now()
	e.Level = level
	e.Format = format
	e.Args = args
	if !e.logger.opt.disableCaller {
		// obtain file name and line number by runtime.Caller().
		// 2 is the stack depth
		if pc, file, line, ok := runtime.Caller(2); !ok {
			e.File = "???"
			e.Func = "???"
		} else {
			e.File, e.Line, e.Func = file, line, runtime.FuncForPC(pc).Name()
			e.Func = e.Func[strings.LastIndex(e.Func, "/")+1:]
		}
	}
	e.format()
	e.writer()
	e.release()
}

func (e *Entry) format() {
	_ = e.logger.opt.formatter.Format(e)
}

func (e *Entry) writer() {
	e.logger.mu.Lock()
	_, _ = e.logger.opt.output.Write(e.Buffer.Bytes())
	e.logger.mu.Unlock()
}

// release can clean entry.Buffer and sync.Pool
func (e *Entry) release() {
	e.Args, e.Line, e.File, e.Format, e.Func = nil, 0, "", "", ""
	e.Buffer.Reset()
	e.logger.entryPool.Put(e)
}

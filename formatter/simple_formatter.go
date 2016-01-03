package formatter

import (
	"bytes"
	"fmt"
	"github.com/zhangyuchen0411/fungolog"
	"time"
)

// it likes `20:46:55.052 [I] Info Message`
func SimpleFormat(buf *bytes.Buffer, level fungolog.Level, args ...interface{}) {
	fmt.Fprint(buf, time.Now().Format("15:04:05.000"), " [", string(level.String()[0]), "] ")
	fmt.Fprint(buf, args...)
}

// it likes `2016-01-03 21:06:42.452 [D] Debug Message`
func SimpleFormatFullTime(buf *bytes.Buffer, level fungolog.Level, args ...interface{}) {
	fmt.Fprint(buf, time.Now().Format("2006-01-02 15:04:05.000"), " [", string(level.String()[0]), "] ")
	fmt.Fprint(buf, args...)
}

// 这个恐怕需要你自己去看了 examples/simple
func SimpleFormatColor(buf *bytes.Buffer, level fungolog.Level, args ...interface{}) {
	var color int
	switch level {
	case fungolog.Info:
		color = CONSOLE_GREEN
	case fungolog.Warning:
		color = CONSOLE_YELLOW
	case fungolog.Error, fungolog.Panic:
		color = CONSOLE_RED
	default:
		color = CONSOLE_WHITE
	}

	fmt.Fprint(buf, time.Now().Format("2006-01-02 15:04:05.000"),
		ConsoleColor(color, CONSOLE_BG_TRANS), " [", string(level.String()[0]), "] ")
	fmt.Fprint(buf, args...)
	buf.WriteString(ConsoleColorReset())
}

// it likes `21:22:29.400 [D] simple.go:35 Debug Message`
// 注意: 通过Write调用时调用者信息会多跳过一层.
func SimpleFormatCallerInfo(buf *bytes.Buffer, level fungolog.Level, args ...interface{}) {
	fmt.Fprint(buf, time.Now().Format("15:04:05.000"), " [", string(level.String()[0]), "] ")
	SimpleCallerInfo(buf, 4)
	buf.WriteByte(' ')
	fmt.Fprint(buf, args...)
}

// 当前文件的路径为全路径, 同样,通过Write调用时调用者信息会多跳过一层.
func SimpleFormatFullCaller(buf *bytes.Buffer, level fungolog.Level, args ...interface{}) {
	fmt.Fprint(buf, time.Now().Format("15:04:05.000"), " [", string(level.String()[0]), "] ")
	FullCallerInfo(buf, 4)
	buf.WriteByte(' ')
	fmt.Fprint(buf, args...)
}

// Error级别及以上时打印出当前goroutine调用堆栈
func SimpleFormatStackInfo(buf *bytes.Buffer, level fungolog.Level, args ...interface{}) {
	SimpleFormatCallerInfo(buf, level, args...)
	if level >= fungolog.Error {
		StackInfo(buf, false)
		buf.WriteString("\r\n\r\n")
	}
}

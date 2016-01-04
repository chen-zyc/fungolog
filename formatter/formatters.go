package formatter

import (
	"bytes"
	"fmt"
	"github.com/zhangyuchen0411/fungolog"
	"path/filepath"
	"runtime"
	"time"
)

func SimpleCallerInfo(buf *bytes.Buffer, skip int) {
	_, filePath, line, ok := runtime.Caller(skip)
	if !ok {
		filePath, line = "???", -1
	} else {
		filePath = filepath.Base(filePath)
	}
	fmt.Fprint(buf, filePath, ":", line)
}

func FullCallerInfo(buf *bytes.Buffer, skip int) {
	_, filePath, line, ok := runtime.Caller(skip)
	if !ok {
		filePath, line = "???", -1
	}
	fmt.Fprint(buf, filePath, ":", line)
}

func StackInfo(buf *bytes.Buffer, all bool) {
	stackBytes := make([]byte, 1024)
	n := runtime.Stack(stackBytes, all)
	buf.Write(stackBytes[:n])
}

func DateTime(buf *bytes.Buffer, layout string) {
	buf.WriteString(time.Now().Format(layout))
}

func SimpleLevel(buf *bytes.Buffer, level fungolog.Level) {
	fmt.Fprint(buf, "[", string(level.String()[0]), "]")
}

func FullLevel(buf *bytes.Buffer, level fungolog.Level) {
	fmt.Fprint(buf, "[", level.String(), "]")
}

func LevelN(buf *bytes.Buffer, level fungolog.Level, n int) {
	fmt.Fprint(buf, '[', level.String()[:n], "]")
}

func Args(buf *bytes.Buffer, args ...interface{}) {
	fmt.Fprint(buf, args...)
}

func ColorByLevel(buf *bytes.Buffer, level fungolog.Level) {
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
	buf.WriteString(ConsoleColor(color, CONSOLE_BG_TRANS))
}

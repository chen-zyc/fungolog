package formatter

import (
	"bytes"
	"fmt"
	"path/filepath"
	"runtime"
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

package main

import (
	"fmt"
	"github.com/zhangyuchen0411/fungolog"
	"github.com/zhangyuchen0411/fungolog/buffers"
	"github.com/zhangyuchen0411/fungolog/formatter"
	"github.com/zhangyuchen0411/fungolog/logger"
	"github.com/zhangyuchen0411/fungolog/writers"
	"sync"
	"time"
)

func main() {
	fmt.Println("++++++++++ SimpleFormatFullTime ++++++++++")
	l := &logger.Logger{
		Format:    formatter.SimpleFormatCallerInfo,
		Threshold: fungolog.Debug,
		Buffers:   buffers.NewDefaultBufferPool(),
		WriteFunc: getWriteFunc(),
	}

	ticker := time.NewTicker(10 * time.Second)
	i := 0
	l.Debug("test message\n")
	for range ticker.C {
		l.Debugln("Debug Message")
		l.Infoln("Info Message")
		l.Warningln("Warning Message")
		l.Errorln("Error Message")
		l.Panicln("Panic Message")
		l.Write(fungolog.Info, "Message from Write()\n")
		i++
		if i > 10 {
			break
		}
	}
	ticker.Stop()
}

// 所有内容统一写到all.log中, Error及以上的信息另写到error.log中, 所有文件都以日期为后缀
func getWriteFunc() func([]byte, fungolog.Level) {
	allFile := writers.NewFileWriter(writers.FileNameGenerator(func() string {
		return "logfiles/all_" + time.Now().Format("200601021504") + ".log"
	}))
	allFileMutex := &sync.Mutex{}
	errFile := writers.NewFileWriter(writers.FileNameGenerator(func() string {
		return "logfiles/error_" + time.Now().Format("200601021504") + ".log"
	}))
	errFileMutex := &sync.Mutex{}

	return func(data []byte, level fungolog.Level) {
		if level >= fungolog.Error {
			errFileMutex.Lock()
			errFile.Write(data)
			errFileMutex.Unlock()
		}
		allFileMutex.Lock()
		allFile.Write(data)
		allFileMutex.Unlock()
	}
}

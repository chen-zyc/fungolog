package logger

import (
	"bytes"
	"fmt"
	"github.com/zhangyuchen0411/fungolog"
	"github.com/zhangyuchen0411/fungolog/buffers"
	"github.com/zhangyuchen0411/fungolog/formatter"
	"os"
	"sync"
)

type Logger struct {
	Format    func(*bytes.Buffer, fungolog.Level, ...interface{})
	WriteFunc func([]byte, fungolog.Level)
	Threshold fungolog.Level
	Buffers   BufferPool
}

type BufferPool interface {
	Put(x *bytes.Buffer)
	Get() *bytes.Buffer
}

func NewLogger(level fungolog.Level) *Logger {
	l := &Logger{
		Threshold: level,
	}
	l.setDefaultFormatter()
	l.setDefaultBufferPool()
	l.setDefaultWriter()
	return l
}

func (this *Logger) Debug(args ...interface{}) {
	this.Write(fungolog.Debug, args...)
}

func (this *Logger) Debugf(format string, args ...interface{}) {
	this.Write(fungolog.Debug, fmt.Sprintf(format, args...))
}

func (this *Logger) Debugln(args ...interface{}) {
	this.Write(fungolog.Debug, fmt.Sprintln(args...))
}

func (this *Logger) Info(args ...interface{}) {
	this.Write(fungolog.Info, args...)
}

func (this *Logger) Infof(format string, args ...interface{}) {
	this.Write(fungolog.Info, fmt.Sprintf(format, args...))
}

func (this *Logger) Infoln(args ...interface{}) {
	this.Write(fungolog.Info, fmt.Sprintln(args...))
}

func (this *Logger) Warning(args ...interface{}) {
	this.Write(fungolog.Warning, args...)
}

func (this *Logger) Warningf(format string, args ...interface{}) {
	this.Write(fungolog.Warning, fmt.Sprintf(format, args...))
}

func (this *Logger) Warningln(args ...interface{}) {
	this.Write(fungolog.Warning, fmt.Sprintln(args...))
}

func (this *Logger) Error(args ...interface{}) {
	this.Write(fungolog.Error, args...)
}

func (this *Logger) Errorf(format string, args ...interface{}) {
	this.Write(fungolog.Error, fmt.Sprintf(format, args...))
}

func (this *Logger) Errorln(args ...interface{}) {
	this.Write(fungolog.Error, fmt.Sprintln(args...))
}

func (this *Logger) Panic(args ...interface{}) {
	this.Write(fungolog.Panic, args...)
}

func (this *Logger) Panicf(format string, args ...interface{}) {
	this.Write(fungolog.Panic, fmt.Sprintf(format, args...))
}

func (this *Logger) Panicln(args ...interface{}) {
	this.Write(fungolog.Panic, fmt.Sprintln(args...))
}

func (this *Logger) Write(level fungolog.Level, args ...interface{}) {
	if level < this.Threshold {
		return
	}
	if this.Buffers == nil {
		this.setDefaultBufferPool()
	}
	if this.Format == nil {
		this.setDefaultFormatter()
	}
	if this.Write == nil {
		this.setDefaultWriter()
	}
	buf := this.Buffers.Get()
	this.Format(buf, level, args...)
	this.WriteFunc(buf.Bytes(), level)
	this.Buffers.Put(buf)
}

func (this *Logger) setDefaultBufferPool() {
	this.Buffers = buffers.NewDefaultBufferPool()
}

func (this *Logger) setDefaultFormatter() {
	this.Format = formatter.SimpleFormat
}

func (this *Logger) setDefaultWriter() {
	m := new(sync.Mutex)
	this.WriteFunc = func(data []byte, level fungolog.Level) {
		m.Lock()
		defer m.Unlock()
		w := os.Stderr
		w.Write(data)
	}
}

var _ fungolog.Logger = &Logger{}

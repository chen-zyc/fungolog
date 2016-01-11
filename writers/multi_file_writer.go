package writers

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/zhangyuchen0411/fungolog"
	"sync"
)

func NewMultiFileWriteFunc(writer *MultiFileWriter) func([]byte, fungolog.Level) {
	if writer == nil {
		return func(data []byte, level fungolog.Level) {}
	}
	writerMutex := sync.Mutex{}
	return func(data []byte, level fungolog.Level) {
		writerMutex.Lock()
		defer writerMutex.Unlock()
		writer.Write(data, level)
	}
}

type MultiFileWriterSetting struct {
	NameGenerator FileNameGenerator
	Async         bool // 是否异步写入
}

// 多文件写入
type MultiFileWriter struct {
	levels map[fungolog.Level]Writer
	others Writer
	all    Writer
}

func NewMultiFileWriter(levels map[fungolog.Level]MultiFileWriterSetting, others MultiFileWriterSetting, all MultiFileWriterSetting) *MultiFileWriter {
	var newWriter = func(setting MultiFileWriterSetting) Writer {
		if setting.NameGenerator == nil {
			return emptyWriter
		}
		if setting.Async {
			return NewAsyncFileWriter(NewFileWriter(setting.NameGenerator), 1024)
		}
		return NewFileWriter(setting.NameGenerator)
	}
	var (
		levelWriters        = make(map[fungolog.Level]Writer, len(levels))
		otherWriter  Writer = newWriter(others)
		allWriter    Writer = newWriter(all)
	)
	for level, setting := range levels {
		levelWriters[level] = newWriter(setting)
	}

	return &MultiFileWriter{
		levels: levelWriters,
		others: otherWriter,
		all:    allWriter,
	}
}

func (this *MultiFileWriter) Write(data []byte, level fungolog.Level) error {
	buf := bytes.Buffer{}
	w := this.levels[level]
	if w != nil {
		err := w.Write(data)
		if err != nil {
			buf.WriteString(fmt.Sprintf("[%s]%s", level.String(), err.Error()))
		}
	} else {
		err := this.others.Write(data)
		if err != nil {
			buf.WriteString(fmt.Sprintf("[OTHER]%s", err.Error()))
		}
	}

	err := this.all.Write(data)
	if err != nil {
		if buf.Len() > 0 {
			buf.WriteString(" | ")
		}
		buf.WriteString(fmt.Sprintf("[ALL]%s", err.Error()))
	}

	if buf.Len() > 0 {
		return errors.New(buf.String())
	}
	return nil
}

func (this *MultiFileWriter) Close() error {
	buf := bytes.Buffer{}
	var err error
	for level, w := range this.levels {
		err = w.Close()
		if err != nil {
			if buf.Len() > 0 {
				buf.WriteString(" | ")
			}
			buf.WriteString(fmt.Sprintf("[%s]%s", level.String(), err.Error()))
		}
	}
	err = this.others.Close()
	if err != nil {
		if buf.Len() > 0 {
			buf.WriteString(" | ")
		}
		buf.WriteString(fmt.Sprintf("[OTHER]%s", err.Error()))
	}
	err = this.all.Close()
	if err != nil {
		if buf.Len() > 0 {
			buf.WriteString(" | ")
		}
		buf.WriteString(fmt.Sprintf("[ALL]%s", err.Error()))
	}
	if buf.Len() > 0 {
		return errors.New(buf.String())
	}
	return nil
}

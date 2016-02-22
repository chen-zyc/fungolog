package writers

import (
	"bytes"
	"testing"
	"time"
)

type bufferWriter struct {
	buf *bytes.Buffer
}

func (w *bufferWriter) Write(data []byte) error {
	_, err := w.buf.Write(data)
	return err
}

func (w *bufferWriter) Close() error {
	return nil
}

func TestAsyncListFileWriter(t *testing.T) {
	writer := &bufferWriter{
		buf: &bytes.Buffer{},
	}
	fileWriter := NewAsyncListFileWriter(writer)
	err := fileWriter.Write([]byte("hello world"))
	if err != nil {
		t.Error(err)
		return
	}
	time.Sleep(5 * time.Second)
	if writer.buf.String() != "hello world" {
		t.Error("not hello world:", writer.buf.String()[:100])
	}
}

func BenchmarkAsyncListFileWriter(b *testing.B) {
	writer := &bufferWriter{
		buf: &bytes.Buffer{},
	}
	fileWriter := NewAsyncListFileWriter(writer)
	for i := 0; i < b.N; i++ {
		err := fileWriter.Write([]byte("hello"))
		if err != nil {
			b.Error(err)
			break
		}
	}
	//	fileWriter.Close()
	b.StopTimer()
	time.Sleep(5 * time.Second)
	b.Log(writer.buf.Len())
}

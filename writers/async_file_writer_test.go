package writers

import (
	"fmt"
	"testing"
	"time"
)

func TestAsyncFileWriter(t *testing.T) {
	w := NewAsyncFileWriter(NewFileWriter(FileNameGenerator(func() string {
		return fmt.Sprintf("test_%s.log", time.Now().Format("200601021504"))
	})), 1024)

	start := time.Now()
	for time.Now().Sub(start) <= 2*time.Minute {
		w.Write([]byte(fmt.Sprintf("log at %s\n", time.Now().Format("2006-01-02 15:04:05.000"))))
		time.Sleep(100 * time.Millisecond)
	}

	w.Close()
}

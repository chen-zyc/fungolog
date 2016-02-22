package writers

import (
	"testing"
	"time"
)

type testWriter struct {
}

func (tw *testWriter) Write(data []byte) error {
	//	time.Sleep(30 * time.Millisecond)
	return nil
}

func (tw *testWriter) Close() error {
	return nil
}

func BenchmarkWrite(b *testing.B) {
	b.StopTimer()
	asyncWriter := NewAsyncFileWriter(&testWriter{}, 1)
	data := []byte("testData")

	// 建几个耗时的goroutine
	for i := 0; i < 100; i++ {
		go func() {
			for {
				for j := 0; j < 10000; j++ {
				}
				time.Sleep(10*time.Millisecond)
			}
		}()
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		asyncWriter.Write(data)
	}
}

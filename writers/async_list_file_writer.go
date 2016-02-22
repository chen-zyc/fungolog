package writers

import (
	"container/list"
	"sync"
	"sync/atomic"
)

type AsyncListFileWriter struct {
	w             Writer
	close         int32
	dataList      *list.List
	dataListMutex sync.Mutex
}

func NewAsyncListFileWriter(w Writer) *AsyncListFileWriter {
	writer := &AsyncListFileWriter{
		w:        w,
		dataList: list.New(),
		close:    0,
	}
	go writer.asyncWrite()
	return writer
}

func (this *AsyncListFileWriter) Write(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	temp := make([]byte, len(data))
	copy(temp, data)
	this.dataListMutex.Lock()
	defer this.dataListMutex.Unlock()
	this.dataList.PushBack(temp)
	return nil
}

func (this *AsyncListFileWriter) Close() error {
	atomic.StoreInt32(&this.close, 1)
	return nil
}

func (this *AsyncListFileWriter) asyncWrite() {
	for {
		if close := atomic.LoadInt32(&this.close); close == 1 {
			break
		}
		data := this.getDataFromList()
		if data == nil {
			continue
		}
		this.w.Write(data)
	}
	this.w.Close()
}

func (this *AsyncListFileWriter) getDataFromList() []byte {
	this.dataListMutex.Lock()
	defer this.dataListMutex.Unlock()
	ele := this.dataList.Front()
	if ele == nil {
		return nil
	}
	this.dataList.Remove(ele)
	if ele.Value == nil {
		return nil
	}
	return ele.Value.([]byte)
}

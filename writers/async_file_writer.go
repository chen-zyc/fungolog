package writers

type AsyncFileWriter struct {
	w     Writer
	c     chan []byte
	close chan bool
}

func NewAsyncFileWriter(w Writer, bufSize int) *AsyncFileWriter {
	aw := &AsyncFileWriter{
		w:     w,
		c:     make(chan []byte, bufSize),
		close: make(chan bool),
	}
	go aw.asyncWrite()
	return aw
}

func (this *AsyncFileWriter) Write(data []byte) error {
	// 如果data是复用的，当asyncWrite()读取时数据可能被复用了
	temp := make([]byte, len(data))
	copy(temp, data)
	this.c <- temp
	return nil
}

func (this *AsyncFileWriter) Close() error {
	this.close <- true
	return nil
}

func (this *AsyncFileWriter) asyncWrite() {
	var data []byte
	for {
		select {
		case <-this.close:
			this.w.Close()
			break
		case data = <-this.c:
			this.w.Write(data)
		}
	}
}

var _ Writer = &AsyncFileWriter{}

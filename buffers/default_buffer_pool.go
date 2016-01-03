package buffers

import (
	"bytes"
	"sync"
)

type DefaultBufferPool struct {
	pool *sync.Pool
}

func NewDefaultBufferPool() *DefaultBufferPool {
	return &DefaultBufferPool{
		pool: &sync.Pool{
			New: func() interface{} {
				return &bytes.Buffer{}
			},
		},
	}
}

func (this *DefaultBufferPool) Get() *bytes.Buffer {
	return this.pool.Get().(*bytes.Buffer)
}

func (this *DefaultBufferPool) Put(x *bytes.Buffer) {
	if x == nil {
		return
	}
	x.Reset()
	this.pool.Put(x)
}

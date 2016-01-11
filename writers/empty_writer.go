package writers

type EmptyWriter struct {
}

var emptyWriter = NewEmptyWriter()

func NewEmptyWriter() *EmptyWriter {
	return &EmptyWriter{}
}

func (this *EmptyWriter) Write(data []byte) error {
	return nil
}

func (this *EmptyWriter) Close() error {
	return nil
}

var _ Writer = &EmptyWriter{}

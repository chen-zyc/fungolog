package writers

type Writer interface {
	Write([]byte) error
	Close() error
}

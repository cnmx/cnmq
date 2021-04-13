package broker

type Queue interface {
	ID() int64
	Next() ([]byte, error)
	Seek() error
}

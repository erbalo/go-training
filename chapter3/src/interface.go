package src

type Reader interface {
	Read() string
}

type Closer interface {
	Closer() error
}

type ReadCloser interface {
	Reader
	Closer
}

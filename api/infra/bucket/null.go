package bucket

import "io"

type Null struct{}

func NewNull() *Null { return &Null{} }

func (s *Null) Get(key string) (io.ReadCloser, error) {
	return nil, nil
}

func (s *Null) Put(key string, r io.Reader) error {
	return nil
}

func (s *Null) URL(key string) string {
	return ""
}

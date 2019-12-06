package bucket

import (
	"bytes"
	"io"
	"io/ioutil"
)

type Mem struct {
	m map[string][]byte
}

type MemFile struct {
	*bytes.Reader
}

func (f *MemFile) Close() error {
	return nil
}

func NewMem() *Mem {
	return &Mem{m: map[string][]byte{}}
}

func (s *Mem) Get(key string) (io.ReadCloser, error) {
	return &MemFile{bytes.NewReader(s.m[key])}, nil
}

func (s *Mem) Put(key string, r io.ReadSeeker) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	s.m[key] = b
	return nil
}

func (s *Mem) URL(key string) string {
	return ""
}

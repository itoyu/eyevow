package bucket

import (
	"fmt"
	"io"
	"os"
	"path"
)

type Local struct {
	root string
	url  string
}

func NewLocal(root, url string) *Local {
	return &Local{root, url}
}

func (s *Local) Get(key string) (io.ReadCloser, error) {
	f, err := os.Open(path.Join(s.root, key))
	return f, err
}

func (s *Local) Put(key string, r io.ReadSeeker) error {
	fn := path.Join(s.root, key)
	dn := path.Dir(fn)

	_, err := os.Lstat(dn)
	if err != nil {
		os.MkdirAll(dn, os.ModePerm)
	}

	f, err := os.Create(fn)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = io.Copy(f, r)
	return err
}

func (s *Local) URL(key string) string {
	return fmt.Sprintf("%s%s", s.url, key)
}

package bucket

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
)

type Bucket interface {
	Get(key string) (io.ReadCloser, error)
	Put(key string, r io.ReadSeeker) error
	URL(key string) string
}

func GetAll(b Bucket, key string) ([]byte, error) {
	f, err := b.Get(key)

	if err != nil {
		return nil, err
	}
	defer f.Close()

	bin, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return bin, nil
}

func PutAll(b Bucket, key string, blob []byte) error {
	buf := bytes.NewReader(blob)
	return b.Put(key, buf)
}

func MarshalJSON(b Bucket, key string, v interface{}) error {
	bin, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return PutAll(b, key, bin)
}

func UnmarshalJSON(b Bucket, key string, v interface{}) error {
	bin, err := GetAll(b, key)
	if err != nil {
		return err
	}
	return json.Unmarshal(bin, v)
}

package redis

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"sync"

	"github.com/klauspost/compress/snappy"

	"sadlil.com/samples/golib/cache"
)

// redisEntry is a wrap around the actual object stored.
type redisEntry struct {
	Compressed bool   `json:"c,omitempty"`
	Data       string `json:"d"`
}

func newRedisEntry(obj interface{}, opt *cache.Option) (*redisEntry, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	e := &redisEntry{Compressed: opt.Compress}
	if opt.Compress {
		data, err = e.compress(data)
		if err != nil {
			return nil, err
		}
		return &redisEntry{Compressed: true, Data: base64.StdEncoding.EncodeToString(data)}, nil
	}
	return &redisEntry{Compressed: false, Data: string(data)}, nil
}

func (e *redisEntry) Unmarshal(obj interface{}) error {
	var d []byte
	var err error

	if e.Compressed {
		var decoded []byte
		decoded, err = base64.StdEncoding.DecodeString(e.Data)
		if err != nil {
			return err
		}

		d, err = e.decompress(decoded)
		if err != nil {
			return err
		}
	} else {
		d = []byte(e.Data)
	}

	err = json.Unmarshal(d, obj)
	if err != nil {
		return err
	}
	return nil
}

// Size ...
func (e *redisEntry) Size() int {
	return len(e.Data)
}

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func (e *redisEntry) compress(b []byte) ([]byte, error) {
	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer func() { bufPool.Put(buf) }()

	sn := snappy.NewBufferedWriter(buf)
	_, err := sn.Write(b)
	if err != nil {
		return nil, err
	}

	if err = sn.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (e *redisEntry) decompress(b []byte) ([]byte, error) {
	sn := snappy.NewReader(bytes.NewReader(b))
	decompressedBytes, err := io.ReadAll(sn)
	if err != nil {
		return nil, err
	}
	return decompressedBytes, nil
}

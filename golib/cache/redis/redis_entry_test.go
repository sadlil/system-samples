package redis

import "testing"

func TestCompressDecompress(t *testing.T) {
	data := `The Snappy compression format in the Go programming language.
	This is a drop-in replacement for github.com/golang/snappy.
	It provides a full, compatible replacement of the Snappy package by simply changing imports.
	See Snappy Compatibility in the S2 documentation.
	"Better" compression mode is used. For buffered streams concurrent compression is used.
	For more options use the s2 package.
	
	usage
	Replace imports github.com/golang/snappy with github.com/klauspost/compress/snappy.`

	e := &redisEntry{}
	cd, err := e.compress([]byte(data))
	if err != nil {
		t.Fatalf("e.compress: got %v, want no error", err)
	}

	// We are hoping the compress data should be smaller
	t.Logf("len(data): got %d, want <= %d", len(cd), len(data))
	if len(data) <= len(cd) {
		t.Errorf("len(data): got %d, want <= %d", len(cd), len(data))
	}

	decompressed, err := e.decompress(cd)
	if err != nil {
		t.Fatalf("e.decompress: got %v, want no error", err)
	}

	// We are hoping the compress data should be smaller
	if data != string(decompressed) {
		t.Errorf("e.decompress: decompressed doesn't match the original data")
	}
}

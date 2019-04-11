package bytes

import (
	"bytes"
	"compress/gzip"
)

/*
	Compress takes a byte array and returns compressed (gzipped) version of that byte array.
*/
func Compress(byt []byte) ([]byte, error) {

	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write(byt); err != nil {
		return nil, err
	}
	if err := gz.Flush(); err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
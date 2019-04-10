package bytes

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
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
	str := base64.StdEncoding.EncodeToString(b.Bytes())
	data, _ := base64.StdEncoding.DecodeString(str)
	rdata := bytes.NewReader(data)
	r, _ := gzip.NewReader(rdata)
	s, _ := ioutil.ReadAll(r)
	return s, nil
}
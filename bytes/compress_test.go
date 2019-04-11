package bytes

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestCompress(t *testing.T) {
	str := "Hello World"
	b := []byte(str)
	compressed, err := Compress(b)
	assert.Nil(t, err)
	assert.NotEqual(t, string(compressed), str)
	assert.Equal(t, "\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xf2H\xcd\xc9\xc9W\b\xcf/\xcaI\x01\x00\x00\x00\xff\xff\x01\x00\x00\xff\xffV\xb1\x17J\v\x00\x00\x00", string(compressed))
	err = ioutil.WriteFile("/tmp/dat1.gz", compressed, 0644)
	assert.Nil(t, err)
	// Manually decompressed /tmp1.gz and it seems fine.
}

package url

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilePath(t *testing.T) {
	path, err := FilePath("https://blog.golang.org/godoc-documenting-go-code")
	assert.Nil(t,err)
	assert.Equal(t,path,"blog.golang.org/godoc-documenting-go-code")
}
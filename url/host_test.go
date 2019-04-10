package url

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHost(t *testing.T) {
	h, err := Host("https://blog.golang.org/godoc-documenting-go-code")
	assert.Nil(t, err)
	assert.Equal(t, "blog.golang.org", h)
}

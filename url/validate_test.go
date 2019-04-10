package url

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidate(t *testing.T) {
	assert.True(t,Validate("https://blog.golang.org/godoc-documenting-go-code"))
	assert.False(t,Validate("htergtps:/\\/blog.golang.org/godoc-documenting-go-code"))
}
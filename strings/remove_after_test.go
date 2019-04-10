package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveAfter(t *testing.T) {
	assert.Equal(t, RemoveAfter("//example.com?abc=123#anchor", "#"), "//example.com?abc=123#")
	assert.Equal(t, RemoveAfter("//example.com?abc=123#*anchor", "#*"), "//example.com?abc=123#*")
}

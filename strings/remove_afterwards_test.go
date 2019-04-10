package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveAfterwards(t *testing.T) {

	assert.Equal(t,RemoveAfterwards("//example.com?abc=123#anchor","#"),"//example.com?abc=123")


	assert.Equal(t,RemoveAfterwards("//example.com?abc=123#*anchor","#*"),"//example.com?abc=123")

}

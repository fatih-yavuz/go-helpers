package function

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func foo() string {
	return Caller()
}


func bar() string {
	return foo()
}


func TestCaller(t *testing.T) {
	c := bar()
	assert.Contains(t,c,"bar")
}

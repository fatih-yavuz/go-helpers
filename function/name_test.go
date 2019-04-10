package function

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func randomFunc() {

}

func TestName(t *testing.T) {

	assert.Contains(t,Name(randomFunc),"randomFunc")

}
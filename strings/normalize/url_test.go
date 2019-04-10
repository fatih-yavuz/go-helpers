package normalize

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUrl(t *testing.T) {

	u, err := Url("www.example.com")
	assert.Nil(t, err)
	assert.Equal(t, "http://www.example.com", u)

	u, err = Url("//example.com?abc=123#anchor")
	assert.Nil(t, err)
	assert.Equal(t, u, "https://example.com?abc=123")

	u, err = Url("src=\"//example.com?abc=123#anchor\"")
	assert.Nil(t, err)
	assert.Equal(t, u, "https://example.com?abc=123")

	u, err = Url("src='//example.com?abc=123#anchor'")
	assert.Nil(t, err)
	assert.Equal(t, u, "https://example.com?abc=123")

	u, err = Url("href='//example.com?abc=123#anchor'")
	assert.Nil(t, err)
	assert.Equal(t, u, "https://example.com?abc=123")

	u, err = Url("href=\"//example.com?abc=123#anchor\"")
	assert.Nil(t, err)
	assert.Equal(t, u, "https://example.com?abc=123")

	_, err = Url("href=\"f*esthabc=123#anchor\"")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "invalid")

}

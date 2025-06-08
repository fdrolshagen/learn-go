package http

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeaders_Add(t *testing.T) {
	headers := Headers{}
	headers.Add("foo", "bar")

	assert.Equal(t, "bar", headers.Get("foo"))
}

func TestHeaders_AddMultipleWithSameKey(t *testing.T) {
	headers := Headers{}
	headers.Add("foo", "bar")
	headers.Add("foo", "baz")

	assert.Equal(t, "bar", headers.Get("foo"))
	assert.Equal(t, []string{"bar", "baz"}, headers.Values("foo"))
}

func TestHeaders_Del(t *testing.T) {
	headers := Headers{}
	headers.Add("foo", "bar")

	assert.Equal(t, "bar", headers.Get("foo"))

	headers.Del("foo")

	assert.Equal(t, "", headers.Get("foo"))
}

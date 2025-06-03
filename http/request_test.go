package http

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseRequest(t *testing.T) {
	input := []byte("" +
		"GET /hello HTTP/1.1\r\n" +
		"Host: example.com\r\n" +
		"User-Agent: TestClient/1.0\r\n" +
		"\r\n" +
		"BodyContentHere")

	req, err := ParseRequest(input)

	assert.Nil(t, err)
	assert.Equal(t, "GET", req.Method)
	assert.Equal(t, "/hello", req.Url)
}

func TestParseRequest_NoHeaderBodySep(t *testing.T) {
	input := []byte("GET / HTTP/1.1\r\nHost: example.com\r\n")

	_, err := ParseRequest(input)

	assert.NotNil(t, err)
}

func TestParseRequest_MalformedRequestLine(t *testing.T) {
	input := []byte("" +
		"\r\n" +
		"\r\n" +
		"BodyContentHere")

	_, err := ParseRequest(input)

	assert.NotNil(t, err)
}

func TestParseRequest_RequestLineIncomplete(t *testing.T) {
	input := []byte("" +
		"GET /hello\r\n" +
		"Host: example.com\r\n" +
		"User-Agent: TestClient/1.0\r\n" +
		"\r\n" +
		"BodyContentHere")

	_, err := ParseRequest(input)

	assert.NotNil(t, err)
}

package http

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseRequest(t *testing.T) {
	input := []byte("" +
		"GET /hello?foo=bar HTTP/1.1\r\n" +
		"Host: example.com\r\n" +
		"User-Agent: TestClient/1.0\r\n" +
		"\r\n" +
		"BODY")

	req, err := ParseRequest(input)

	assert.Nil(t, err)
	assert.Equal(t, "GET", req.Method)
	assert.Equal(t, "/hello", req.Url)

	assert.Len(t, req.QueryParams, 1)
	assert.Equal(t, "bar", req.QueryParams["foo"])

	assert.Equal(t, "HTTP/1.1", req.Protocol)
	assert.Equal(t, 1, req.ProtocolMajor)
	assert.Equal(t, 1, req.ProtocolMinor)

	assert.Len(t, req.Headers, 2)
	assert.Equal(t, "example.com", req.Headers["Host"])
	assert.Equal(t, "TestClient/1.0", req.Headers["User-Agent"])

	assert.Equal(t, []byte("BODY"), req.Body)
}

func TestParseRequest_NoHeaderBodySep(t *testing.T) {
	input := []byte("GET / HTTP/1.1\r\nHost: example.com\r\nBODY")

	_, err := ParseRequest(input)

	assert.NotNil(t, err)
}

func TestParseRequest_ProtocolInvalid(t *testing.T) {
	input := []byte("GET / 1.1\r\nHost: example.com\r\n")

	_, err := ParseRequest(input)

	assert.NotNil(t, err)
}

func TestParseRequest_ProtocolVersionIncomplete(t *testing.T) {
	input := []byte("GET / HTTP/1\r\nHost: example.com\r\n")

	_, err := ParseRequest(input)

	assert.NotNil(t, err)
}

func TestParseRequest_ProtocolVersionInvalid(t *testing.T) {
	input := []byte("GET / HTTP/1.a\r\nHost: example.com\r\n")

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

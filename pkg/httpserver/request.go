package httpserver

import (
	"bytes"
	"fmt"
)

type HttpRequest struct {
	url      string
	method   string
	protocol string
	headers  []string
	body     []byte
}

func ParseHttpRequest(b []byte) (HttpRequest, error) {
	var httpRequest HttpRequest

	separator := []byte("\r\n\r\n")
	idx := bytes.Index(b, separator)
	if idx < 0 {
		return httpRequest, fmt.Errorf("malformed http request: no header/body separator")
	}

	headerBlock := b[:idx]
	httpRequest.body = b[idx+len(separator):]

	headerLines := bytes.Split(headerBlock, []byte("\r\n"))
	if len(headerLines) < 1 {
		return httpRequest, fmt.Errorf("malformed http request: no request line")
	}

	requestLine := bytes.Split(headerLines[0], []byte(" "))
	if len(requestLine) != 3 {
		return httpRequest, fmt.Errorf("malformed http request: request line incomplete")
	}

	httpRequest.method = string(requestLine[0])
	httpRequest.url = string(requestLine[1])
	httpRequest.protocol = string(requestLine[2])

	// TODO parse headers, not needed yet

	return httpRequest, nil
}

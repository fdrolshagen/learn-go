package httpserver

import (
	"bytes"
	"fmt"
)

type HttpRequest struct {
	Url           string
	Params        map[string]string
	Method        string
	Protocol      string
	ProtocolMajor int
	ProtocolMinor int
	headers       map[string]string
	Body          []byte
}

func ParseHttpRequest(b []byte) (HttpRequest, error) {
	var httpRequest HttpRequest

	separator := []byte("\r\n\r\n")
	idx := bytes.Index(b, separator)
	if idx < 0 {
		return httpRequest, fmt.Errorf("malformed http request: no header/body separator")
	}

	headerBlock := b[:idx]
	httpRequest.Body = b[idx+len(separator):]

	headerLines := bytes.Split(headerBlock, []byte("\r\n"))
	if len(headerLines) < 1 {
		return httpRequest, fmt.Errorf("malformed http request: no request line")
	}

	requestLine := bytes.Split(headerLines[0], []byte(" "))
	if len(requestLine) != 3 {
		return httpRequest, fmt.Errorf("malformed http request: request line incomplete")
	}

	httpRequest.Method = string(requestLine[0])
	httpRequest.Url = string(requestLine[1])
	httpRequest.Protocol = string(requestLine[2])

	// TODO parse headers, not needed yet

	return httpRequest, nil
}

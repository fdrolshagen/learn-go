package http

import (
	"bytes"
	"fmt"
	"strings"
)

// Request holds all information about a valid HTTP-Request
type Request struct {
	Url           string
	Params        map[string]string
	Method        string
	Protocol      string
	ProtocolMajor int
	ProtocolMinor int
	Headers       map[string]string
	Body          []byte
}

// ParseRequest parses a []byte and tries to construct a valid Request.
// Returns an error if the []byte cannot be parsed as a valid HTTP-Request
func ParseRequest(b []byte) (Request, error) {
	var req Request

	separator := []byte("\r\n\r\n")
	idx := bytes.Index(b, separator)
	if idx < 0 {
		return req, fmt.Errorf("malformed http request: no header/body separator")
	}

	headerBlock := b[:idx]
	req.Body = b[idx+len(separator):]

	headerLines := bytes.Split(headerBlock, []byte("\r\n"))
	if len(headerLines) < 1 {
		return req, fmt.Errorf("malformed http request: no request line")
	}

	requestLine := bytes.Split(headerLines[0], []byte(" "))
	if len(requestLine) != 3 {
		return req, fmt.Errorf("malformed http request: request line incomplete")
	}

	req.Method = string(requestLine[0])
	req.Url = string(requestLine[1])
	req.Protocol = string(requestLine[2])

	for _, h := range headerLines[1:] {
		if req.Headers == nil {
			req.Headers = map[string]string{}
		}
		header := strings.SplitN(string(h), ":", 2)
		req.Headers[strings.TrimSpace(header[0])] = strings.TrimSpace(header[1])
	}

	// TODO parse query params

	return req, nil
}

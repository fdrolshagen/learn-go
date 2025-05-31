package http

import (
	"bytes"
	"fmt"
	"path"
	"strings"
)

// Request holds all information about a valid HTTP-Request
type Request struct {
	Url           string
	QueryParams   map[string]string
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

	// requestLine contains: GET /url HTTP1.1
	requestLine := bytes.Split(headerLines[0], []byte(" "))
	if len(requestLine) != 3 {
		return req, fmt.Errorf("malformed http request: request line incomplete")
	}

	req.Method = string(requestLine[0])
	req.Protocol = string(requestLine[2])

	url := string(requestLine[1])
	url = path.Clean(url)
	req.Url = url

	// url might contain: /url?key=value&foo=bar
	urlSplitted := strings.Split(url, "?")
	if len(urlSplitted) > 2 {
		return req, fmt.Errorf("malformed http request: url contains more than one '?': %s", url)
	}

	if len(urlSplitted) == 2 {
		req.Url = urlSplitted[0]
		req.QueryParams = map[string]string{}
		params := strings.Split(urlSplitted[1], "&")
		for _, param := range params {
			kv := strings.Split(param, "=")
			if len(kv) == 2 {
				req.QueryParams[kv[0]] = kv[1]
			}
		}
	}

	for _, h := range headerLines[1:] {
		if req.Headers == nil {
			req.Headers = map[string]string{}
		}
		header := strings.SplitN(string(h), ":", 2)
		req.Headers[strings.TrimSpace(header[0])] = strings.TrimSpace(header[1])
	}

	return req, nil
}

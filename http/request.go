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
func ParseRequest(b []byte) (req Request, err error) {
	header, body, err := splitHeaderAndBody(b)
	if err != nil {
		return req, err
	}
	req.Body = body

	headerLines, err := splitHeader(header)
	if err != nil {
		return req, err
	}

	if err := parseRequestLine(headerLines[0], &req); err != nil {
		return req, err
	}

	if err := parseHeaders(headerLines[1:], &req); err != nil {
		return req, err
	}

	return req, nil
}

func splitHeaderAndBody(b []byte) ([]byte, []byte, error) {
	separator := []byte("\r\n\r\n")
	idx := bytes.Index(b, separator)
	if idx < 0 {
		return nil, nil, fmt.Errorf("malformed http request: no header/body separator")
	}
	return b[:idx], b[idx+len(separator):], nil
}

func splitHeader(header []byte) ([][]byte, error) {
	headerLines := bytes.Split(header, []byte("\r\n"))
	if len(headerLines) < 1 {
		return nil, fmt.Errorf("malformed http request: no request line")
	}
	return headerLines, nil
}

func parseRequestLine(line []byte, req *Request) error {
	parts := bytes.Split(line, []byte(" "))
	if len(parts) != 3 {
		return fmt.Errorf("malformed http request: request line incomplete")
	}
	req.Method = string(parts[0])
	req.Protocol = string(parts[2])

	return parseURL(string(parts[1]), req)
}

func parseURL(rawUrl string, req *Request) error {
	url := path.Clean(rawUrl)
	req.Url = url

	// url might contain: /url?key=value&foo=bar
	urlSplitted := strings.Split(url, "?")
	if len(urlSplitted) > 2 {
		return fmt.Errorf("malformed http request: url contains more than one '?': %s", req.Url)
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

	return nil
}

func parseHeaders(headerLines [][]byte, req *Request) error {
	if len(headerLines) == 0 {
		return nil
	}

	req.Headers = make(map[string]string)
	for _, h := range headerLines[1:] {
		header := strings.SplitN(string(h), ":", 2)
		if len(header) != 2 {
			return fmt.Errorf("malformed http request: header incomplete")
		}

		key := strings.TrimSpace(header[0])
		if key == "" {
			return fmt.Errorf("malformed http request: empty header key")
		}

		req.Headers[key] = strings.TrimSpace(header[1])
	}

	return nil
}

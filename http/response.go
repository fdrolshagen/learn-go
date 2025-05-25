package http

import (
	"fmt"
	"strings"
	"time"
)

// Response holds all information about a HTTP-Response
type Response struct {
	StatusCode  int
	Body        string
	ContentType string
}

// RawResponse converts the Response to a valid HTTP-Response containing all necessary
// protocol headers and body
func (r *Response) RawResponse() (string, error) {
	var b strings.Builder

	if _, err := fmt.Fprintf(&b, "HTTP/1.1 %d %s\r\n", r.StatusCode, StatusCodes[r.StatusCode]); err != nil {
		return "", err
	}

	if _, err := fmt.Fprintf(&b, "Date: %s\r\n", getHttpDate()); err != nil {
		return "", err
	}

	if _, err := fmt.Fprintf(&b, "Server: GO-HTTP (0.1)\r\n"); err != nil {
		return "", err
	}

	if r.Body != "" {
		if _, err := fmt.Fprintf(&b, "Content-Length: %d\r\n", len(r.Body)); err != nil {
			return "", err
		}
	}

	if r.ContentType != "" {
		if _, err := fmt.Fprintf(&b, "Content-type: %s\r\n", r.ContentType); err != nil {
			return "", err
		}
	}

	if _, err := fmt.Fprintf(&b, "\r\n"); err != nil {
		return "", err
	}

	if _, err := fmt.Fprintf(&b, "%s", r.Body); err != nil {
		return "", err
	}

	return b.String(), nil
}

func getHttpDate() string {
	const format = "Mon, 02 Jan 2006 15:04:05 GMT"
	return time.Now().UTC().Format(format)
}

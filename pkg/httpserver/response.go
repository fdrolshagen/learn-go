package httpserver

import (
	"fmt"
	"strings"
	"time"
)

type HttpResponse struct {
	statusCode  int
	body        string
	contentType string
}

func (r *HttpResponse) RawHttpResponse() (string, error) {
	var b strings.Builder

	if _, err := fmt.Fprintf(&b, "HTTP/1.1 %d %s\r\n", r.statusCode, StatusCodes[r.statusCode]); err != nil {
		return "", err
	}

	if _, err := fmt.Fprintf(&b, "Date: %s\r\n", getHttpDate()); err != nil {
		return "", err
	}

	if _, err := fmt.Fprintf(&b, "Server: GO-HTTP (0.1)\r\n"); err != nil {
		return "", err
	}

	if r.body != "" {
		if _, err := fmt.Fprintf(&b, "Content-Length: %d\r\n", len(r.body)); err != nil {
			return "", err
		}
	}

	if r.contentType != "" {
		if _, err := fmt.Fprintf(&b, "Content-type: %s\r\n", r.contentType); err != nil {
			return "", err
		}
	}

	if _, err := fmt.Fprintf(&b, "\r\n"); err != nil {
		return "", err
	}

	if _, err := fmt.Fprintf(&b, "%s", r.body); err != nil {
		return "", err
	}

	return b.String(), nil
}

func getHttpDate() string {
	const format = "Mon, 02 Jan 2006 15:04:05 GMT"
	return time.Now().UTC().Format(format)
}

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

const (
	HTTPVersion    = "HTTP/1.1"
	ServerName     = "GO-HTTP (0.1)"
	HTTPDateFormat = "Mon, 02 Jan 2006 15:04:05 GMT"
	CRLF           = "\r\n"
)

// RawResponse converts the Response to a valid HTTP-Response containing all necessary
// protocol headers and body
func (r *Response) RawResponse() (string, error) {
	var b strings.Builder

	// Write status line
	if _, err := fmt.Fprintf(&b, "%s %d %s%s", HTTPVersion, r.StatusCode, StatusCodes[r.StatusCode], CRLF); err != nil {
		return "", err
	}

	// Write mandatory headers
	if err := writeHeader(&b, "Date", date()); err != nil {
		return "", fmt.Errorf("failed to write status line: %w", err)
	}

	if err := writeHeader(&b, "Server", ServerName); err != nil {
		return "", fmt.Errorf("failed to write server header: %w", err)
	}

	// Write optional headers
	if r.Body != "" {
		if err := writeHeader(&b, "Content-Length", fmt.Sprint(len(r.Body))); err != nil {
			return "", fmt.Errorf("failed to write content length: %w", err)
		}

	}

	if r.ContentType != "" {
		if err := writeHeader(&b, "Content-Type", r.ContentType); err != nil {
			return "", fmt.Errorf("failed to write content type: %w", err)
		}
	}

	// Write Body
	if _, err := fmt.Fprintf(&b, "%s%s", CRLF, r.Body); err != nil {
		return "", fmt.Errorf("failed to write body: %w", err)
	}

	return b.String(), nil
}

func writeHeader(b *strings.Builder, name, value string) error {
	_, err := fmt.Fprintf(b, "%s: %s%s", name, value, CRLF)
	return err
}

func date() string {
	return time.Now().UTC().Format(HTTPDateFormat)
}

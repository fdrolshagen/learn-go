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

func CreateRawResponse(httpResponse HttpResponse) string {
	var b strings.Builder

	fmt.Fprintf(&b, "HTTP/1.1 %d %s\r\n", httpResponse.statusCode, StatusCodes[httpResponse.statusCode])
	fmt.Fprintf(&b, "Date: %s\r\n", getHttpDate())
	fmt.Fprintf(&b, "Server: GO-HTTP (0.1)\r\n")

	if httpResponse.contentType != "" {
		fmt.Fprintf(&b, "Content-type: %s\r\n", httpResponse.contentType)
	}

	fmt.Fprintf(&b, "\r\n")
	fmt.Fprintf(&b, "%s", httpResponse.body)

	return b.String()
}

func getHttpDate() string {
	const httpTimeformat = "Mon, 02 Jan 2006 15:04:05 GMT"
	return time.Now().UTC().Format(httpTimeformat)

}

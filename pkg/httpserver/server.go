package httpserver

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

var PORT int = 9999

var StatusCodes map[int]string = map[int]string{
	200: "ok",
	404: "not found",
}

type HttpRequest struct {
	url      string
	method   string
	protocol string
	headers  []string
	body     []byte
}

type HttpResponse struct {
	statusCode  int
	body        string
	contentType string
}

type Config struct {
	Port      string
	StaticDir string
}

type Server struct {
	Port      string
	StaticDir string
}

func CreateServer(config Config) *Server {
	return &Server{
		Port:      config.Port,
		StaticDir: config.StaticDir,
	}
}

// TODO make the server configurable: port, static dir

func (s *Server) StartServer() {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(PORT))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	log.Printf("Now listening on port %d\n", PORT)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		if err == io.EOF {
			return
		}
		log.Printf("Cannot read Request %s", err)
		return
	}

	httpRequest, err := parseHttpRequest(buf)
	if err != nil {
		log.Printf("Cannot parse HttpRequest %s", err)
		return
	}

	var httpResponse HttpResponse

	var requestedPath string
	if httpRequest.url == "/" {
		requestedPath = "/index.html"
	} else {
		requestedPath = httpRequest.url
	}

	httpResponse.body, err = readFile(s.StaticDir + requestedPath)
	if err != nil {
		httpResponse.statusCode = 404
	} else {
		httpResponse.contentType = determineContentType(requestedPath)
		httpResponse.statusCode = 200
	}

	log.Printf("Incoming request: %s %s -> %d", httpRequest.method, httpRequest.url, httpResponse.statusCode)

	response := buildRawResponse(httpResponse)
	conn.Write([]byte(response))
}

func readFile(filePath string) (string, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(file), nil
}

func determineContentType(url string) string {
	switch ext := path.Ext(url); ext {
	case ".html":
		return "text/html"
	case ".json":
		return "application/json"
	default:
		return "text/plain"
	}
}

func parseHttpRequest(b []byte) (HttpRequest, error) {
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

func buildRawResponse(httpResponse HttpResponse) string {
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

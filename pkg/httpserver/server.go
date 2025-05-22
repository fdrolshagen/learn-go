package httpserver

import (
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

type Config struct {
	Port      int
	StaticDir string
}

type Server struct {
	Port      int
	StaticDir string
}

func CreateServer(config Config) *Server {
	return &Server{
		Port:      config.Port,
		StaticDir: config.StaticDir,
	}
}

func (s *Server) StartServer() {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(s.Port))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Now listening on port %d\n", s.Port)

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

	httpRequest, err := ParseHttp(buf)
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
		httpResponse.contentType = GuessContentType(requestedPath)
		httpResponse.statusCode = 200
	}

	log.Printf("Incoming request: %s %s -> %d", httpRequest.method, httpRequest.url, httpResponse.statusCode)

	response := CreateRawResponse(httpResponse)
	_, err = conn.Write([]byte(response))
}

func readFile(filePath string) (string, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(file), nil
}

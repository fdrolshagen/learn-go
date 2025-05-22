package httpserver

import (
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

type Config struct {
	Port      int
	StaticDir string
}

type Server struct {
	Port      int
	StaticDir string
	Router    *Router
}

func CreateServer(config Config, router *Router) *Server {
	return &Server{
		Port:      config.Port,
		StaticDir: config.StaticDir,
		Router:    router,
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

	request, err := ParseHttpRequest(buf)
	if err != nil {
		log.Printf("Cannot parse HttpRequest %s", err)
		return
	}

	route := s.Router.selectRoute(request.method, request.url)
	originalUrl := request.url
	request.url = rewriteUrl(request.url, route.path)
	response, err := route.handler.handle(request)
	if err != nil {
		response = HttpResponse{}
		response.statusCode = 500
	}

	log.Printf("Incoming request: %s %s -> %d", request.method, originalUrl, response.statusCode)

	raw, err := response.RawHttpResponse()
	if err != nil {
		return
	}

	// TODO implement chunking
	_, err = conn.Write([]byte(raw))
}

func rewriteUrl(url string, prefix string) string {
	if prefix == "/" {
		return url
	}
	return strings.TrimPrefix(url, prefix)
}

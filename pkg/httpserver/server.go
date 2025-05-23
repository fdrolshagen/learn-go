package httpserver

import (
	"io"
	"log"
	"net"
	"strconv"
)

type Server struct {
	Port      int
	StaticDir string
	Router    *Router
}

func CreateServer(port int, router *Router) *Server {
	return &Server{
		Port:   port,
		Router: router,
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
		log.Printf("Cannot parse Request %s", err)
		return
	}

	route := s.Router.selectRoute(request.Method, request.Url)
	handler := PanicRecoveryMiddleware(route.handler)
	handler = RewriteAfterRoutingMiddleware(handler, route.path)
	handler = LoggingMiddleware(handler)

	response, err := handler.Handle(request)
	if err != nil {
		response = Response{
			StatusCode:  500,
			ContentType: "text/plain",
			Body:        "internal server error",
		}
	}

	raw, err := response.RawHttpResponse()
	if err != nil {
		return
	}

	// TODO implement chunking
	_, err = conn.Write([]byte(raw))
}

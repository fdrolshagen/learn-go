package http

import (
	"io"
	"log"
	"net"
	"strconv"
)

type Server struct {
	Port   int
	Router *Router
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

	log.Printf("Now listening on %s\n", ln.Addr().String())

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

	request, err := ParseRequest(buf)
	if err != nil {
		log.Printf("Cannot parse Request %s", err)
		return
	}

	route := s.Router.selectRoute(request.Method, request.Url)
	resp := s.handle(route, request)

	raw, err := resp.RawResponse()
	if err != nil {
		return
	}

	// TODO implement chunking
	_, err = conn.Write([]byte(raw))
}

func (s *Server) handle(route Route, req Request) Response {
	handle := PanicRecoveryMiddleware(route.handle)
	handle = RewriteAfterRoutingMiddleware(handle, route.path)
	handle = LoggingMiddleware(handle)

	response, err := handle(req)
	if err != nil {
		return Response{
			StatusCode:  500,
			ContentType: "text/plain",
			Body:        "internal server error",
		}
	}

	return response
}

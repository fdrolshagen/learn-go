package http

import (
	"fmt"
	"log"
	"net"
	"runtime/debug"
	"strconv"
	"strings"
)

const (
	MaxClientRequestSize = 1024 * 1024
)

type Server struct {
	Port        int
	Router      *Router
	mr          *MetricRegistry
	middlewares []Middleware
}

func CreateServer(port int, router *Router) *Server {
	mr := CreateMetricRegistry()
	ar := mr.getActuatorRoute()
	router.prependRoute(ar)

	return &Server{
		Port:   port,
		Router: router,
		mr:     mr,
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

	buf := make([]byte, MaxClientRequestSize)
	err := read(conn, buf)
	if err != nil {
		log.Printf("error reading request: %s", err)
		return
	}

	// just a prototype for now
	go s.mr.Increment("total_request_count")

	resp := s.process(buf)
	raw, err := resp.RawResponse()
	if err != nil {
		return
	}

	// TODO implement chunking
	_, err = conn.Write([]byte(raw))
}

func read(conn net.Conn, buf []byte) error {
	_, err := conn.Read(buf)
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}

	return nil
}

func (s *Server) process(buf []byte) Response {
	request, err := ParseRequest(buf)
	if err != nil {
		log.Printf("Cannot parse Request %s", err)
		resp, _ := HandleInternalServerError(request)
		return resp
	}

	route := s.Router.selectRoute(request.Method, request.Url)
	rewrite(&request, route.path)
	return recoverHandle(s.handle, route, request)
}

func (s *Server) handle(route Route, req Request) Response {
	handle := DefaultAccessLogMiddleware(route.handle)

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

func rewrite(req *Request, prefix string) {
	p := req.Url

	if prefix != "/" {
		p = strings.TrimPrefix(p, prefix)
	}

	if p == "" {
		p = "/"
	}

	req.Url = p
}

func recoverHandle(f func(Route, Request) Response, route Route, req Request) (resp Response) {
	defer func(req Request) {
		if r := recover(); r != nil {
			log.Printf("PANIC: %v\n%s", r, debug.Stack())
			resp, _ = HandleInternalServerError(req)
		}
	}(req)

	return f(route, req)
}

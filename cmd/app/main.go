package main

import (
	"fdrolshagen/server/pkg/httpserver"
	"flag"
)

func main() {
	port := flag.Int("port", 8080, "Port to listen on")
	dir := flag.String("dir", "./", "Directory to serve")
	flag.Parse()

	staticHandler := httpserver.CreateStaticHandler(*dir)
	router := httpserver.CreateRouter()

	router.GET("/web", staticHandler)
	router.GET("/api", ApiHandler{})
	server := httpserver.CreateServer(*port, router)
	server.StartServer()
}

type ApiHandler struct{}

func (h ApiHandler) Handle(httpserver.Request) (httpserver.Response, error) {
	panic("PANIC!")
}

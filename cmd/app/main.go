package main

import (
	"fdrolshagen/server/pkg/http"
	"flag"
)

func main() {
	port := flag.Int("port", 8080, "Port to listen on")
	dir := flag.String("dir", "./", "Directory to serve")
	flag.Parse()

	staticHandler := http.CreateStaticHandler(*dir)
	router := http.CreateRouter()

	router.GET("/web", staticHandler)
	router.GET("/api", ApiHandler{})
	server := http.CreateServer(*port, router)
	server.StartServer()
}

type ApiHandler struct{}

func (h ApiHandler) Handle(http.Request) (http.Response, error) {
	panic("PANIC!")
}

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

	router.AddRoute("GET", "/web", staticHandler)
	router.AddRoute("GET", "/api", ApiHandler{})
	server := httpserver.CreateServer(*port, router)
	server.StartServer()
}

type ApiHandler struct{}

func (h ApiHandler) Handle(request httpserver.HttpRequest) (httpserver.HttpResponse, error) {
	var response httpserver.HttpResponse
	response.StatusCode = 200
	response.ContentType = "application/json"
	response.Body = "{\"key\":\"value\"}"
	return response, nil
}

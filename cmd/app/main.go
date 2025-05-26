package main

import (
	"fdrolshagen/learn-go/http"
	"flag"
)

func main() {
	port := flag.Int("port", 8080, "Port to listen on")
	dir := flag.String("dir", "./", "Directory to serve")
	flag.Parse()

	router := http.CreateRouter()
	router.WithMiddleware(http.DefaultAccessLogMiddleware)

	staticHandler := http.StaticHandler{StaticDir: *dir}
	router.GET("/web", staticHandler.Handle)
	router.GET("/health", Health)

	server := http.CreateServer(*port, router)
	server.StartServer()
}

func Health(http.Request) (http.Response, error) {
	return http.Response{
		StatusCode:  200,
		ContentType: "application/json",
		Body:        "{\"status\": \"up\"}",
	}, nil
}

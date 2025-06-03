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
	router.WithMiddleware(SecurityMiddleware)
	router.GET("/health", Health)
	router.Mount("/", *dir)

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

func SecurityMiddleware(next http.Handle) http.Handle {
	return func(req http.Request) (resp http.Response, err error) {
		// check Authorization header or do something else before/after further processing
		return next(req)
	}
}

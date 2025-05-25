
# Learn GO

This repository contains a personal collection of packages implemented for the purpose of learning the GO programming language.

## HTTP

Minimal implementation of HTTP 1.1 (according to [RFC 2616](https://datatracker.ietf.org/doc/html/rfc2616)).

This package allows the user to create a simple HTTP server with routing capabilities and middleware support.

[!IMPORTANT] OBVIOUSLY: DO NOT USE THIS IN PRODUCTION! ;)

### Usage

```go
package main

import (
	"flag"
	"github.com/fdrolshagen/learn-go/http"
)

func main() {
	port := flag.Int("port", 8080, "Port to listen on")
	dir := flag.String("dir", "./", "Directory to serve")
	flag.Parse()

	staticHandler := http.StaticHandler{StaticDir: *dir}
	router := http.CreateRouter()

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
```
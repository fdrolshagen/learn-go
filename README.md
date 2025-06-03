
# Learn GO

This repository contains a personal collection of packages implemented for the purpose of learning the GO programming language.

## HTTP

Minimal implementation of HTTP 1.1 (according to [RFC 2616](https://datatracker.ietf.org/doc/html/rfc2616)) without using net/http from the standard library.

The following features are currently implemented and might be extended depending on my spare time :laughing:
- create and start an HTTP server
- routing with user defined handlers
- provied a default implementation for a static handler to serve files from disk
- middleware support to easily extend the functionality

:exclamation: OBVIOUSLY: DO NOT USE THIS IN PRODUCTION! :smirk:

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
		// check Authorization header or something else before/after further processing
		return next(req)
	}
}
```

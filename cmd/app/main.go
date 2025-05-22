package main

import (
	"fdrolshagen/server/pkg/httpserver"
)

func main() {
	config := httpserver.Config{
		Port: 9999,
	}

	staticHandler := httpserver.CreateStaticHandler("./web/static")
	router := httpserver.CreateRouter()

	router.AddRoute("GET", "/foobar", staticHandler)
	server := httpserver.CreateServer(config, router)
	server.StartServer()
}

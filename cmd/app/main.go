package main

import (
	"fdrolshagen/server/pkg/httpserver"
)

func main() {
	config := httpserver.Config{
		Port:      9999,
		StaticDir: "./web/static",
	}

	server := httpserver.CreateServer(config)
	server.StartServer()
}

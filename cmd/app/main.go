package main

import (
	"fdrolshagen/server/pkg/http"
	"flag"
)

func main() {
	port := flag.Int("port", 8080, "Port to listen on")
	dir := flag.String("dir", "./", "Directory to serve")
	flag.Parse()

	staticHandler := http.StaticHandler{StaticDir: *dir}
	router := http.CreateRouter()

	router.GET("/web", staticHandler.Handle)
	router.GET("/panic", HandleWithPanic)
	server := http.CreateServer(*port, router)
	server.StartServer()
}

func HandleWithPanic(http.Request) (http.Response, error) {
	panic("PANIC!")
}

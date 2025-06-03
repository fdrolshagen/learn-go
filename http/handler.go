package http

import (
	"os"
	"path"
)

// Handle needs to be implemented to be passed as a Handler function for a Route
type Handle func(Request) (Response, error)

// HandleNotFound is a default implementation which always returns an HTTP 404 (not found) response
func HandleNotFound(_ Request) (Response, error) {
	return Response{
		StatusCode: 404,
	}, nil
}

// StaticHandler holds information about the static directory used for the default implementation
type StaticHandler struct {
	StaticDir string
}

// Handle is a default implementation inside the StaticHandler which allows to serve a directory
// on disk
func (h StaticHandler) Handle(request Request) (Response, error) {
	var resp Response

	var requestedPath string
	if request.Url == "/" {
		requestedPath = "/index.html"
	} else {
		requestedPath = request.Url
	}

	file, err := readFile(h.StaticDir + requestedPath)
	if err != nil {
		resp.StatusCode = 404
	} else {
		resp.ContentType = GuessContentType(requestedPath)
		resp.StatusCode = 200
		resp.Body = file
	}

	return resp, nil
}

func readFile(filePath string) (string, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(file), nil
}

func GuessContentType(url string) string {
	switch ext := path.Ext(url); ext {
	case ".html":
		return "text/html"
	case ".json":
		return "application/json"
	default:
		return "text/plain"
	}
}

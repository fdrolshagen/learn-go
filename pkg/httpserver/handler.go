package httpserver

import "os"

type Handler interface {
	Handle(HttpRequest) (HttpResponse, error)
}

type NotFoundHandler struct {
}

func (n NotFoundHandler) Handle(request HttpRequest) (HttpResponse, error) {
	return HttpResponse{
		StatusCode: 404,
	}, nil
}

type StaticHandler struct {
	StaticDir string
}

func CreateStaticHandler(dir string) StaticHandler {
	return StaticHandler{
		StaticDir: dir,
	}
}

func (h StaticHandler) Handle(request HttpRequest) (HttpResponse, error) {
	var httpResponse HttpResponse

	var requestedPath string
	if request.Url == "/" {
		requestedPath = "/index.html"
	} else {
		requestedPath = request.Url
	}

	file, err := readFile(h.StaticDir + requestedPath)
	if err != nil {
		httpResponse.StatusCode = 404
	} else {
		httpResponse.ContentType = GuessContentType(requestedPath)
		httpResponse.StatusCode = 200
		httpResponse.Body = file
	}

	return httpResponse, nil
}

func readFile(filePath string) (string, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(file), nil
}

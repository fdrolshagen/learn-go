package http

import "os"

type Handler interface {
	Handle(Request) (Response, error)
}

type NotFoundHandler struct {
}

func (n NotFoundHandler) Handle(request Request) (Response, error) {
	return Response{
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

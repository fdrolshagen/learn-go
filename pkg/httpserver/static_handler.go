package httpserver

import "os"

func HandleStatic(request HttpRequest, staticDir string) HttpResponse {
	var httpResponse HttpResponse

	var requestedPath string
	if request.url == "/" {
		requestedPath = "/index.html"
	} else {
		requestedPath = request.url
	}

	file, err := readFile(staticDir + requestedPath)
	if err != nil {
		httpResponse.statusCode = 404
	} else {
		httpResponse.contentType = GuessContentType(requestedPath)
		httpResponse.statusCode = 200
		httpResponse.body = file
	}

	return httpResponse
}

func readFile(filePath string) (string, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(file), nil
}

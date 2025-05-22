package httpserver

type NotFoundHandler struct {
}

func (n NotFoundHandler) handle(request HttpRequest) (HttpResponse, error) {
	var response HttpResponse
	response.statusCode = 404
	return response, nil
}

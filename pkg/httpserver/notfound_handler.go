package httpserver

type NotFoundHandler struct {
}

func (n NotFoundHandler) Handle(request HttpRequest) (HttpResponse, error) {
	var response HttpResponse
	response.StatusCode = 404
	return response, nil
}

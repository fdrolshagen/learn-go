package httpserver

type NotFoundHandler struct {
}

func (n NotFoundHandler) Handle(request HttpRequest) (HttpResponse, error) {
	return HttpResponse{
		StatusCode: 404,
	}, nil
}

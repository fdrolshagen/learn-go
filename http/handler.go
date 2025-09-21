package http

// HandleFunc is a function type that processes HTTP requests and returns responses.
// It takes a Request parameter and returns a Response and an error.
// Implementations of this type serve as request handlers in the HTTP server,
// allowing custom processing of incoming HTTP requests through a Router
type HandleFunc func(Request) (Response, error)

// HandleNotFound implements the HandleFunc type and serves as a default implementation
// when an HTTP 404 should be returned.
func HandleNotFound(_ Request) (Response, error) {
	return Response{
		StatusCode:  404,
		Body:        StatusCodes[404],
		ContentType: "text/plain",
	}, nil
}

func HandleRequestTooLarge(_ Request) (Response, error) {
	return Response{
		StatusCode:  413,
		Body:        StatusCodes[413],
		ContentType: "text/plain",
	}, nil
}

func HandleInternalServerError(_ Request) (Response, error) {
	return Response{
		StatusCode:  500,
		Body:        StatusCodes[500],
		ContentType: "text/plain",
	}, nil
}

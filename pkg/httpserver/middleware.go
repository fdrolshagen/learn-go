package httpserver

import (
	"log"
	"runtime/debug"
	"strings"
)

type Middleware func(next Handler) Handler

type HandlerFunc func(req HttpRequest) (HttpResponse, error)

func (f HandlerFunc) Handle(req HttpRequest) (HttpResponse, error) {
	return f(req)
}

func PanicRecoveryMiddleware(next Handler) Handler {
	return HandlerFunc(func(req HttpRequest) (resp HttpResponse, err error) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("PANIC: %v\n%s", r, debug.Stack())
				resp = HttpResponse{
					StatusCode:  500,
					Body:        "Internal Server Error",
					ContentType: "text/plain; charset=utf-8",
				}
				err = nil
			}
		}()

		return next.Handle(req)
	})
}

func RewriteAfterRoutingMiddleware(next Handler, prefix string) Handler {
	return HandlerFunc(func(req HttpRequest) (resp HttpResponse, err error) {
		if prefix != "/" {
			req.Url = strings.TrimPrefix(req.Url, prefix)
		}
		return next.Handle(req)
	})
}

func LoggingMiddleware(next Handler) Handler {
	return HandlerFunc(func(req HttpRequest) (resp HttpResponse, err error) {
		originalUrl := req.Url
		resp, err = next.Handle(req)
		log.Printf("Incoming request: %s %s -> %d", req.Method, originalUrl, resp.StatusCode)
		return resp, nil
	})
}

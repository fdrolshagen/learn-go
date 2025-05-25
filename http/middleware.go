package http

import (
	"log"
	"runtime/debug"
	"strings"
)

type Middleware func(next Handle) Handle

func PanicRecoveryMiddleware(next Handle) Handle {
	return func(req Request) (resp Response, err error) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("PANIC: %v\n%s", r, debug.Stack())
				resp = Response{
					StatusCode:  500,
					Body:        "Internal Server Error",
					ContentType: "text/plain; charset=utf-8",
				}
				err = nil
			}
		}()

		return next(req)
	}
}

func RewriteAfterRoutingMiddleware(next Handle, prefix string) Handle {
	return func(req Request) (resp Response, err error) {
		if prefix != "/" {
			req.Url = strings.TrimPrefix(req.Url, prefix)
		}
		return next(req)
	}
}

func LoggingMiddleware(next Handle) Handle {
	return func(req Request) (resp Response, err error) {
		originalUrl := req.Url
		resp, err = next(req)
		log.Printf("Incoming request: %s %s -> %d", req.Method, originalUrl, resp.StatusCode)
		return resp, nil
	}
}

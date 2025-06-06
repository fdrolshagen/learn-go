package http

import (
	"log"
	"runtime/debug"
	"strings"
	"time"
)

type Middleware func(next HandleFunc) HandleFunc

func PanicRecoveryMiddleware(next HandleFunc) HandleFunc {
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

func RewriteAfterRoutingMiddleware(next HandleFunc, prefix string) HandleFunc {
	return func(req Request) (resp Response, err error) {
		p := req.Url

		if prefix != "/" {
			p = strings.TrimPrefix(p, prefix)
		}

		if p == "" {
			p = "/"
		}

		req.Url = p
		return next(req)
	}
}

func DefaultAccessLogMiddleware(next HandleFunc) HandleFunc {
	return func(req Request) (resp Response, err error) {
		start := time.Now()
		url := req.Url

		resp, err = next(req)

		log.Printf("Incoming request: %s %s --> %d (%d ms)", req.Method, url, resp.StatusCode, time.Since(start).Microseconds())
		return resp, err
	}
}

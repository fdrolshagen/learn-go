package httpserver

import (
	"strings"
)

type Router struct {
	routes []Route
}

type Route struct {
	method  string
	path    string
	handler Handler
}

type Handler interface {
	Handle(HttpRequest) (HttpResponse, error)
}

func CreateRouter() *Router {
	return &Router{}
}

func (r *Router) AddRoute(method string, path string, handler Handler) {
	r.routes = append(r.routes, Route{method: method, path: path, handler: handler})
}

func (r *Router) selectRoute(method string, path string) Route {
	for _, route := range r.routes {
		if route.method == method && strings.HasPrefix(path, route.path) {
			return route
		}
	}

	return Route{method: method, path: path, handler: NotFoundHandler{}}
}

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

func CreateRouter() *Router {
	return &Router{}
}

func (r *Router) GET(path string, handler Handler) {
	r.addRoute("GET", path, handler)
}

func (r *Router) POST(path string, handler Handler) {
	r.addRoute("POST", path, handler)
}

func (r *Router) PUT(path string, handler Handler) {
	r.addRoute("PUT", path, handler)
}

func (r *Router) DELETE(path string, handler Handler) {
	r.addRoute("DELETE", path, handler)
}

func (r *Router) HEAD(path string, handler Handler) {
	r.addRoute("HEAD", path, handler)
}

func (r *Router) addRoute(method string, path string, handler Handler) {
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

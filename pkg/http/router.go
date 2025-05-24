package http

import (
	"strings"
)

type Router struct {
	routes []Route
}

type Route struct {
	method string
	path   string
	handle Handle
}

func CreateRouter() *Router {
	return &Router{}
}

func (r *Router) GET(path string, handle Handle) {
	r.addRoute(GET, path, handle)
}

func (r *Router) POST(path string, handle Handle) {
	r.addRoute(POST, path, handle)
}

func (r *Router) PUT(path string, handle Handle) {
	r.addRoute(PUT, path, handle)
}

func (r *Router) DELETE(path string, handle Handle) {
	r.addRoute(DELETE, path, handle)
}

func (r *Router) HEAD(path string, handle Handle) {
	r.addRoute(HEAD, path, handle)
}

func (r *Router) PATCH(path string, handle Handle) {
	r.addRoute(PATCH, path, handle)
}

func (r *Router) addRoute(method string, path string, handle Handle) {
	r.routes = append(r.routes, Route{method: method, path: path, handle: handle})
}

func (r *Router) selectRoute(method string, path string) Route {
	for _, route := range r.routes {
		if route.method == method && strings.HasPrefix(path, route.path) {
			return route
		}
	}

	return Route{method: method, path: path, handle: HandleNotFound}
}

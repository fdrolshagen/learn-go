package http

import (
	"strings"
)

// Router holds all added routes
type Router struct {
	routes []Route
}

// Route holds information about a single route
type Route struct {
	method string
	path   string
	handle Handle
}

// CreateRouter Creates and Returns a Router instance
func CreateRouter() *Router {
	return &Router{}
}

// GET add a Handle for the specified path using the "GET" method
func (r *Router) GET(path string, handle Handle) {
	r.addRoute(GET, path, handle)
}

// POST add a Handle for the specified path using the "POST" method
func (r *Router) POST(path string, handle Handle) {
	r.addRoute(POST, path, handle)
}

// PUT add a Handle for the specified path using the "PUT" method
func (r *Router) PUT(path string, handle Handle) {
	r.addRoute(PUT, path, handle)
}

// DELETE add a Handle for the specified path using the "DELETE" method
func (r *Router) DELETE(path string, handle Handle) {
	r.addRoute(DELETE, path, handle)
}

// HEAD add a Handle for the specified path using the "HEAD" method
func (r *Router) HEAD(path string, handle Handle) {
	r.addRoute(HEAD, path, handle)
}

// PATCH add a Handle for the specified path using the "PATCH" method
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

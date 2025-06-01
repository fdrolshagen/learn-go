package http

import "strings"

// Router holds all added routes
type Router struct {
	routes      []Route
	middlewares []Middleware
}

// Route holds information about a single route
type Route struct {
	method string
	path   string
	handle Handle
	static bool
}

// CreateRouter Creates and Returns a Router instance
func CreateRouter() *Router {
	return &Router{}
}

// GET add a Handle for the specified path using the "GET" method
func (r *Router) GET(path string, handle Handle) {
	r.addRoute(GET, path, r.stackHandles(handle))
}

// POST add a Handle for the specified path using the "POST" method
func (r *Router) POST(path string, handle Handle) {
	r.addRoute(POST, path, r.stackHandles(handle))
}

// PUT add a Handle for the specified path using the "PUT" method
func (r *Router) PUT(path string, handle Handle) {
	r.addRoute(PUT, path, r.stackHandles(handle))
}

// DELETE add a Handle for the specified path using the "DELETE" method
func (r *Router) DELETE(path string, handle Handle) {
	r.addRoute(DELETE, path, r.stackHandles(handle))
}

// HEAD add a Handle for the specified path using the "HEAD" method
func (r *Router) HEAD(path string, handle Handle) {
	r.addRoute(HEAD, path, r.stackHandles(handle))
}

// PATCH add a Handle for the specified path using the "PATCH" method
func (r *Router) PATCH(path string, handle Handle) {
	r.addRoute(PATCH, path, r.stackHandles(handle))
}

// WithStatic adds a Handle for the specified path using the "GET" method
// Route-Selection allows nested path sections
// example: /web/ and /web/index.html will both match
func (r *Router) WithStatic(path string, handle Handle) {
	r.addStaticRoute(GET, path, r.stackHandles(handle))
}

// WithMiddleware configures a Middleware to be used on all routes within this Router.
// Middleware always needs to be configured before adding a Route
func (r *Router) WithMiddleware(middleware Middleware) {
	r.middlewares = append(r.middlewares, middleware)
}

func (r *Router) addRoute(method string, path string, handle Handle) {
	r.routes = append(r.routes, Route{method: method, path: path, handle: handle, static: false})
}

func (r *Router) addStaticRoute(method string, path string, handle Handle) {
	r.routes = append(r.routes, Route{method: method, path: path, handle: handle, static: true})
}

func (r *Router) selectRoute(method string, path string) Route {
	for _, route := range r.routes {
		// add support for path params like: /users/{user}
		// will need to parse these somewhere and pass to the handler
		if route.method == method && pathMatches(route, path) {
			return route
		}
	}

	return Route{method: method, path: path, handle: HandleNotFound}
}

func pathMatches(route Route, path string) bool {
	if route.static {
		return strings.HasPrefix(path, route.path)
	}

	return path == route.path
}

func (r *Router) stackHandles(handle Handle) Handle {
	var stackedHandle = handle
	for idx := range r.middlewares {
		stackedHandle = r.middlewares[idx](stackedHandle)
	}
	return stackedHandle
}

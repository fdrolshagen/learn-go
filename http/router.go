package http

import "strings"

// Router holds all added routes
type Router struct {
	routes      []Route
	middlewares []Middleware
}

// Route holds information about a single route
type Route struct {
	method  string
	path    string
	handle  HandleFunc
	isMount bool
}

// CreateRouter Creates and Returns a Router instance
func CreateRouter() *Router {
	return &Router{}
}

// GET add a HandleFunc for the specified path using the "GET" method
func (r *Router) GET(path string, handle HandleFunc) {
	r.addRoute(GET, path, r.stackHandles(handle))
}

// POST add a HandleFunc for the specified path using the "POST" method
func (r *Router) POST(path string, handle HandleFunc) {
	r.addRoute(POST, path, r.stackHandles(handle))
}

// PUT add a HandleFunc for the specified path using the "PUT" method
func (r *Router) PUT(path string, handle HandleFunc) {
	r.addRoute(PUT, path, r.stackHandles(handle))
}

// DELETE add a HandleFunc for the specified path using the "DELETE" method
func (r *Router) DELETE(path string, handle HandleFunc) {
	r.addRoute(DELETE, path, r.stackHandles(handle))
}

// HEAD add a HandleFunc for the specified path using the "HEAD" method
func (r *Router) HEAD(path string, handle HandleFunc) {
	r.addRoute(HEAD, path, r.stackHandles(handle))
}

// PATCH add a HandleFunc for the specified path using the "PATCH" method
func (r *Router) PATCH(path string, handle HandleFunc) {
	r.addRoute(PATCH, path, r.stackHandles(handle))
}

// Mount adds a StaticHandler serving the contents of dir on the specified path.
// Paths that should have routing precedence, should be configured before mounting
func (r *Router) Mount(path string, dir string) {
	h := CreateStaticHandler(dir)
	r.addStaticRoute(GET, path, r.stackHandles(h.Handle))
}

// WithMiddleware configures a Middleware to be used on all routes within this Router.
// Middleware always needs to be configured before adding a Route
func (r *Router) WithMiddleware(middleware Middleware) {
	r.middlewares = append(r.middlewares, middleware)
}

func (r *Router) addRoute(method string, path string, handle HandleFunc) {
	r.routes = append(r.routes, Route{method: method, path: path, handle: handle, isMount: false})
}

func (r *Router) addStaticRoute(method string, path string, handle HandleFunc) {
	r.routes = append(r.routes, Route{method: method, path: path, handle: handle, isMount: true})
}

func (r *Router) prependRoute(route Route) {
	routes := append([]Route{route}, r.routes...)
	r.routes = routes
}

func (r *Router) selectRoute(method string, path string) Route {
	for _, route := range r.routes {
		// add support for path params like: /users/{user}
		// will need to parse these somewhere and pass to the handler
		if route.method == method && pathMatches(route, path) {
			return route
		}
	}

	return Route{method: method, path: path, handle: HandleNotFound, isMount: false}
}

func (r *Router) stackHandles(handle HandleFunc) HandleFunc {
	var stackedHandle = handle
	for idx := range r.middlewares {
		stackedHandle = r.middlewares[idx](stackedHandle)
	}
	return stackedHandle
}

func pathMatches(route Route, path string) bool {
	if route.isMount {
		return strings.HasPrefix(path, route.path)
	}

	return path == route.path
}

package http

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	router := CreateRouter()

	assert.NotNil(t, router)
}

func TestRouter_GET(t *testing.T) {
	router := CreateRouter()
	router.GET("/foo", NOPHandler)

	assert.Len(t, router.routes, 1)
	assert.Equal(t, GET, router.routes[0].method)
	assert.Equal(t, "/foo", router.routes[0].path)
}

func TestRouter_PUT(t *testing.T) {
	router := CreateRouter()
	router.PUT("/foo", NOPHandler)

	assert.Len(t, router.routes, 1)
	assert.Equal(t, PUT, router.routes[0].method)
	assert.Equal(t, "/foo", router.routes[0].path)
}

func TestRouter_DELETE(t *testing.T) {
	router := CreateRouter()
	router.DELETE("/foo", NOPHandler)

	assert.Len(t, router.routes, 1)
	assert.Equal(t, DELETE, router.routes[0].method)
	assert.Equal(t, "/foo", router.routes[0].path)
}

func TestRouter_POST(t *testing.T) {
	router := CreateRouter()
	router.POST("/foo", NOPHandler)

	assert.Len(t, router.routes, 1)
	assert.Equal(t, POST, router.routes[0].method)
	assert.Equal(t, "/foo", router.routes[0].path)
}

func TestRouter_HEAD(t *testing.T) {
	router := CreateRouter()
	router.HEAD("/foo", NOPHandler)

	assert.Len(t, router.routes, 1)
	assert.Equal(t, HEAD, router.routes[0].method)
	assert.Equal(t, "/foo", router.routes[0].path)
}

func TestRouter_PATCH(t *testing.T) {
	router := CreateRouter()
	router.PATCH("/foo", NOPHandler)

	assert.Len(t, router.routes, 1)
	assert.Equal(t, PATCH, router.routes[0].method)
	assert.Equal(t, "/foo", router.routes[0].path)
}

func TestRouter_WithMiddleware(t *testing.T) {
	router := CreateRouter()
	router.WithMiddleware(NOPMiddleware)

	assert.Len(t, router.middlewares, 1)
}

func TestRouter_selectRoute_routeIsConfigured(t *testing.T) {
	router := CreateRouter()
	router.GET("/foo", NOPHandler)
	route := router.selectRoute(GET, "/foo")

	assert.Equal(t, ptr(NOPHandler), ptr(route.handle))
}

func TestRouter_selectRoute_pathIsNotConfigured(t *testing.T) {
	router := CreateRouter()
	router.GET("/foo", NOPHandler)
	route := router.selectRoute(GET, "/bar")

	assert.Equal(t, ptr(HandleNotFound), ptr(route.handle))
}

func TestRouter_selectRoute_methodIsNotConfigured(t *testing.T) {
	router := CreateRouter()
	router.GET("/foo", NOPHandler)
	route := router.selectRoute(POST, "/foo")

	assert.Equal(t, ptr(HandleNotFound), ptr(route.handle))
}

func TestRouter_selectRoute_routeIsSubRoute(t *testing.T) {
	router := CreateRouter()
	router.GET("/foo", NOPHandler)
	route := router.selectRoute("GET", "/foo/bar")

	assert.Equal(t, ptr(HandleNotFound), ptr(route.handle), "Sub Route should not match and return HandleNotFound")
}

func TestRouter_selectRoute_routeIsStatic(t *testing.T) {
	router := CreateRouter()
	router.Mount("/foo", "./")
	route := router.selectRoute("GET", "/foo/bar")

	assert.True(t, route.isMount)
}

func NOPHandler(_ Request) (Response, error) {
	return Response{}, nil
}

func NOPMiddleware(next HandleFunc) HandleFunc {
	return func(req Request) (resp Response, err error) {
		return next(req)

	}
}

func ptr(f any) uintptr {
	return reflect.ValueOf(f).Pointer()
}

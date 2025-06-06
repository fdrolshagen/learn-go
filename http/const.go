package http

var StatusCodes = map[int]string{
	200: "ok",
	201: "created",
	202: "accepted",
	204: "no content",
	400: "bad_request",
	401: "unauthorized",
	403: "forbidden",
	404: "not found",
	405: "method not allowed",
	406: "not acceptable",
	409: "conflict",
	500: "internal server error",
	501: "not implemented",
	502: "bad gateway",
}

const (
	GET     = "GET"
	POST    = "POST"
	PUT     = "PUT"
	DELETE  = "DELETE"
	PATCH   = "PATCH"
	HEAD    = "HEAD"
	OPTIONS = "OPTIONS"
)

const (
	APPLICATION_JSON = "application/json"
	TEXT_HTML        = "text/html"
	TEXT_PLAIN       = "text/plain"
	TEXT_XML         = "text/xml"
	IMAGE_PNG        = "image/png"
	IMAGE_JPEG       = "image/jpeg"
	IMAGE_GIF        = "image/gif"
	IMAGE_SVG        = "image/svg+xml"
	APPLICATION_PDF  = "application/pdf"
)

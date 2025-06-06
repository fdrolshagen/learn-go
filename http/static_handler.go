package http

import "path"

// StaticHandler holds information about the static directory used for the default implementation
type StaticHandler struct {
	Reader FileReader
	Dir    string
}

func CreateStaticHandler(dir string) *StaticHandler {
	return &StaticHandler{Reader: OSReader{}, Dir: dir}
}

// Handle serves static files from the local filesystem directory specified in StaticHandler.StaticDir.
// It implements the HandleFunc type for static file serving.
func (h StaticHandler) Handle(req Request) (resp Response, err error) {
	p := getPath(&req)
	file, err := h.Reader.ReadFile(h.Dir + p)
	if err != nil {
		return HandleNotFound(req)
	}

	resp.ContentType = GuessContentType(p)
	resp.StatusCode = 200
	resp.Body = file

	return resp, nil
}

func getPath(req *Request) string {
	if req.Url == "/" {
		return "/index.html"
	}
	return req.Url
}

func GuessContentType(url string) string {
	switch ext := path.Ext(url); ext {
	case ".html":
		return TEXT_HTML
	case ".json":
		return APPLICATION_JSON
	case ".pdf":
		return APPLICATION_PDF
	default:
		return TEXT_PLAIN
	}
}

package http

import "path"

func GuessContentType(url string) string {
	switch ext := path.Ext(url); ext {
	case ".html":
		return "text/html"
	case ".json":
		return "application/json"
	default:
		return "text/plain"
	}
}

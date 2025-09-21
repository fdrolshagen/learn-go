package http

import (
	"encoding/base64"
	"strings"
)

type BasicAuth struct {
	Username string
	Password string
}

func (b BasicAuth) Handler(next HandleFunc) HandleFunc {
	return func(req Request) (resp Response, err error) {
		authHeader := req.Headers.Get("Authorization")
		if authHeader == "" {
			return unauthorized(), nil
		}

		value := strings.SplitN(authHeader, " ", 2)
		if len(value) != 2 && value[0] != "Basic" {
			return unauthorized(), nil
		}

		decoded, err := base64.StdEncoding.DecodeString(value[1])
		if err != nil {
			return unauthorized(), nil
		}

		credentials := strings.SplitN(string(decoded), ":", 2)
		if err != nil {
			return unauthorized(), nil
		}

		username := strings.TrimSpace(credentials[0])
		password := strings.TrimSpace(credentials[1])
		if username != b.Username || password != b.Password {
			return unauthorized(), nil
		}

		return next(req)
	}
}

func unauthorized() Response {
	headers := make(Headers)
	headers.Add("WWW-Authenticate", `Basic realm="Restricted"`)

	return Response{
		StatusCode: 401,
		Headers:    headers,
	}
}

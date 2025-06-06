package http

import "os"

type FileReader interface {
	ReadFile(p string) (string, error)
}

type OSReader struct{}

func (OSReader) ReadFile(p string) (string, error) {
	file, err := os.ReadFile(p)
	if err != nil {
		return "", err
	}

	return string(file), nil
}

package http

type Headers map[string][]string

func (h Headers) Add(key, value string) {
	if _, ok := h[key]; !ok {
		h[key] = []string{value}
		return
	}

	h[key] = append(h[key], value)
}

func (h Headers) Get(key string) string {
	if val, ok := h[key]; ok {
		return val[0]
	}
	return ""
}

func (h Headers) Values(key string) []string {
	return h[key]
}

func (h Headers) Del(key string) {
	delete(h, key)
}

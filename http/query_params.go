package http

type QueryParams map[string]string

func (q QueryParams) Add(key, value string) {
	q[key] = value
}

func (q QueryParams) Get(key string) string {
	if val, ok := q[key]; ok {
		return val
	}
	return ""
}

func (q QueryParams) Del(key string) {
	delete(q, key)
}

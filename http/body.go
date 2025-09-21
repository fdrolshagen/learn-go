package http

import "encoding/json"

type Body []byte

func (b Body) String() string {
	return string(b)
}

func (b Body) Bytes() []byte {
	return b
}

func (b Body) Len() int {
	return len(b)
}

func (b Body) IsEmpty() bool {
	return len(b) == 0
}

func (b Body) Json() (map[string]interface{}, error) {
	if b.IsEmpty() {
		return nil, nil
	}

	var result map[string]interface{}
	err := json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (b Body) Unmarshal(v interface{}) error {
	if b.IsEmpty() {
		return nil
	}
	return json.Unmarshal(b, v)
}

package web

import (
	"encoding/json"
	"io"
)

func DecodePostBody[T any](body io.Reader) (T, error) {
	var parser T

	err := json.NewDecoder(body).Decode(&parser)

	return parser, err
}

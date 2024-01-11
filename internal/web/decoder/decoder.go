package web

import (
	"encoding/json"
	"io"
	"log/slog"
)

func DecodePostBody[T any](body io.Reader) (T, error) {
	var parser T

	err := json.NewDecoder(body).Decode(&parser)

	if err != nil {
		slog.Error("error when decode post body", err)
	}

	return parser, err
}

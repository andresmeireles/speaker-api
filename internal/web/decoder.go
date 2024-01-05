package web

import (
	"encoding/json"
	"io"

	"github.com/andresmeireles/speaker/internal/logger"
)

func DecodePostBody[T any](body io.Reader) (T, error) {
	var parser T

	err := json.NewDecoder(body).Decode(&parser)

	if err != nil {
		logger.Error(err)
	}

	return parser, err
}

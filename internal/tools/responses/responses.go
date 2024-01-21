package responses

import (
	"log/slog"
	"net/http"
)

func BadResponse(w http.ResponseWriter, err error) {
	slog.Error("Error on request process", "error", err)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Error on request"))
}

func DecodeError(w http.ResponseWriter, err error) {
	slog.Error("Error when decode body", "error", err)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Wrong or bad formed request"))
}

func Unauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("Unauthorized"))
}

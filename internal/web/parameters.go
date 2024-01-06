package web

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func badFormattedUrl(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("bad formatted url"))
}

func GetParameter(
	r *http.Request,
	name string,
) (string, error, func(w http.ResponseWriter)) {
	if name == "" {
		return "", fmt.Errorf("no parameter"), badFormattedUrl
	}

	parameter := chi.URLParam(r, name)

	return parameter, nil, nil
}

func GetIntParameter(
	r *http.Request,
	name string,
) (int, error, func(w http.ResponseWriter)) {
	parameter, err, errorFunc := GetParameter(r, name)
	if err != nil {
		slog.Error("error on decode", err)

		return 0, err, errorFunc
	}

	intParameter, err := strconv.Atoi(parameter)
	if err != nil {
		slog.Error("error on decode", err)

		return 0, err, badFormattedUrl
	}

	return intParameter, nil, nil
}

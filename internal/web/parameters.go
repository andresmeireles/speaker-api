package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func badFormattedUrl(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("bad formatted url"))
}

func GetParameter(r *http.Request, name string) string {
	if name == "" {
		panic("cannot search empty parameter")
	}

	parameter := chi.URLParam(r, name)

	if parameter == "" {
		panic("no parameter " + name + " in " + r.URL.Path)
	}

	return parameter
}

func GetIntParameter(r *http.Request, name string) int {
	parameter := GetParameter(r, name)
	intParameter, err := strconv.Atoi(parameter)

	if err != nil {
		panic(fmt.Sprintf("failed to parse parameter %s in %s", name, r.URL.Path))
	}

	return intParameter
}

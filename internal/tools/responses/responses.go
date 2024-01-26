package responses

import (
	"fmt"
	"log/slog"
	"net/http"
	"runtime"
	"strings"

	"github.com/andresmeireles/speaker/internal/tools/env"
)

const (
	BYTE        = 1024
	TRACES_BACK = 3
)

// Create a 400 response
func BadResponse(w http.ResponseWriter, err error) {
	errorLog(err)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Error on request"))
}

// Create a 400 response when body cannot be decoded
func DecodeError(w http.ResponseWriter, err error) {
	errorLog(err)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Error when decode body"))
}

// Create a 401 response
func Unauthorized(w http.ResponseWriter, err error) {
	errorLog(err)
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("Unauthorized"))
}

func stack() []string {
	stack := make([]byte, BYTE)
	runtime.Stack(stack, false)
	strStack := string(stack)

	return strings.Split(strStack, "\n")
}

func whereErrorWasTriggered() []string {
	pc, file, line, ok := runtime.Caller(TRACES_BACK)
	if !ok {
		return []string{"unknown"}
	}

	funcName := runtime.FuncForPC(pc).Name()

	return []string{"function", funcName, "file", file, "line", fmt.Sprint(line)}
}

func errorLog(err error) {
	args := []any{}

	if env.IsDev() {
		for _, v := range whereErrorWasTriggered() {
			args = append(args, v)
		}
	}

	if env.IsDev() && env.ShowStackTrace() {
		for _, v := range stack() {
			args = append(args, v)
		}
	}

	slog.Error(err.Error(), args...)
}

// Create a 200 response
func Ok(w http.ResponseWriter, message []byte) {
	success(w, http.StatusOK, message)
}

// Create a 202 response, ideal for updating put requests
func Accepted(w http.ResponseWriter, message []byte) {
	success(w, http.StatusAccepted, message)
}

// Create a 201 response, ideal for post request
func Created(w http.ResponseWriter, message []byte) {
	success(w, http.StatusCreated, message)
}

// Create a 204 response, ideal for delete
func NoContent(w http.ResponseWriter, message []byte) {
	success(w, http.StatusNoContent, message)
}

func success(w http.ResponseWriter, status int, message []byte) {
	w.WriteHeader(status)
	w.Write(message)
}

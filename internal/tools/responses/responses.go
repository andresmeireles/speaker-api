package responses

import (
	"log/slog"
	"net/http"
)

// Create a 400 response
func BadResponse(w http.ResponseWriter, err error) {
	slog.Error("Error on request process", "error", err)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Error on request"))
}

// Create a 400 response when body cannot be decoded
func DecodeError(w http.ResponseWriter, err error) {
	slog.Error("Error when decode body", "error", err)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Wrong or bad formed request"))
}

// Create a 401 response
func Unauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("Unauthorized"))
}

// Create a 200 response
func Ok(w http.ResponseWriter, message string) {
	success(w, http.StatusOK, message)
}

// Create a 202 response, ideal for updating put requests
func Accepted(w http.ResponseWriter, message string) {
	success(w, http.StatusAccepted, message)
}

// Create a 201 response, ideal for post request
func Created(w http.ResponseWriter, message string) {
	success(w, http.StatusCreated, message)
}

// Create a 204 response, ideal for delete
func NoContent(w http.ResponseWriter, message string) {
	success(w, http.StatusNoContent, message)
}

func success(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	w.Write([]byte(message))
}

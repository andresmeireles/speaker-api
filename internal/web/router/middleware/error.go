package middleware

import (
	"log/slog"
	"net/http"
	"runtime"
)

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				stack := make([]byte, 4*1024)
				runtime.Stack(stack, true)

				slog.Error("panic error", err, string(stack))
				http.Error(w, "internal error, please try again", http.StatusInternalServerError)

				return
			}
		}()

		next.ServeHTTP(w, r)
	})
}

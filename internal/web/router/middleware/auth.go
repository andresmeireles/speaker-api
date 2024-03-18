package middleware

import (
	"log/slog"
	"net/http"

	"github.com/andresmeireles/speaker/internal/auth"
	"github.com/andresmeireles/speaker/internal/tools/responses"
)

func GetTokenFromRequest(req *http.Request) (string, error) {
	if auth := req.Header.Get("Authorization"); auth != "" {
		a := auth[7:]

		return a, nil
	}

	cookie, err := req.Cookie("session_id")
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}

// check if user cookie is valid, if not check if has authorization token and
// check if is valid.
func CheckTokenOnCookie(next http.Handler, authActions auth.Service) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		token, err := GetTokenFromRequest(request)
		if err != nil {
			slog.Error("error on cookie or auth", "cookie", err)
			response.WriteHeader(http.StatusBadRequest)
			response.Write([]byte("Token not found"))

			return
		}
		if err := authActions.ValidateJwt(token); err != nil {
			responses.Unauthorized(response, err)

			return
		}

		next.ServeHTTP(response, request)
	})
}

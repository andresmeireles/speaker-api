package middleware

import (
	"net/http"

	"github.com/andresmeireles/speaker/internal/modules/auth"
)

func CheckTokenOnCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		cookie, err := request.Cookie("token")

		if err != nil {
			response.WriteHeader(http.StatusBadRequest)
			response.Write([]byte("Token not found"))
			return
		}

		if ok := auth.ValidateJwt(cookie.Value); !ok {
			response.WriteHeader(http.StatusUnauthorized)
			response.Write([]byte("Unauthorized"))
			return
		}

		next.ServeHTTP(response, request)
	})
}

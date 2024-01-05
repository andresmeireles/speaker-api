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

		authEntity, err := auth.AuthRepository{}.GetByHash(cookie.Value)

		if err != nil {
			unauthorized(response)
			return
		}

		if authEntity.Expired {
			unauthorized(response)
			return
		}

		if ok := auth.ValidateJwt(authEntity.Hash); !ok {
			auth.ExpireAuth(authEntity, auth.AuthRepository{})
			unauthorized(response)
			return
		}

		next.ServeHTTP(response, request)
	})
}

func unauthorized(res http.ResponseWriter) {
	res.WriteHeader(http.StatusUnauthorized)
	res.Write([]byte("Unauthorized"))
}
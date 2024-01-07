package middleware

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/andresmeireles/speaker/internal/modules/auth"
)

func CheckTokenOnCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		cookie, err := request.Cookie("session_id")
		if err != nil {
			slog.Error("error on cookie", "cookie", err)
			response.WriteHeader(http.StatusBadRequest)
			response.Write([]byte("Token not found"))

			return
		}
		authActions := auth.NewActions()
		authEntity, err := auth.AuthRepository{}.GetByHash(cookie.Value)
		if err != nil {
			slog.Error("error on repository", "cookie", err)
			unauthorized(response)

			return
		}
		if authEntity.Expired {
			slog.Error("auth expired", "cookie", err)
			unauthorized(response)

			return
		}
		if ok := authActions.ValidateJwt(authEntity.Hash); !ok {
			slog.Error("auth expired")
			authActions.Logout(authEntity.UserId)
			unauthorized(response)

			return
		}

		ctx := context.WithValue(request.Context(), "user_id", authEntity.UserId)
		next.ServeHTTP(response, request.WithContext(ctx))
	})
}

func unauthorized(res http.ResponseWriter) {
	res.WriteHeader(http.StatusUnauthorized)
	res.Write([]byte("Unauthorized"))
}

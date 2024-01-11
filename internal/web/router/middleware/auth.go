package middleware

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/andresmeireles/speaker/internal/modules/auth"
	"github.com/andresmeireles/speaker/internal/tools/servicelocator"
)

func getToken(req *http.Request) (string, error) {
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
func CheckTokenOnCookie(next http.Handler, sl servicelocator.ServiceLocator) http.Handler {
	authActions := servicelocator.Get[auth.Actions](sl)

	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		token, err := getToken(request)
		if err != nil {
			slog.Error("error on cookie or auth", "cookie", err)
			response.WriteHeader(http.StatusBadRequest)
			response.Write([]byte("Token not found"))

			return
		}
		authEntity, err := auth.AuthRepository{}.GetByHash(token)
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

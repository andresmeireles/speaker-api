package middleware

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/andresmeireles/speaker/internal/auth"
	"github.com/andresmeireles/speaker/internal/tools/responses"
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
			responses.Unauthorized(response)

			return
		}
		if authEntity.Expired {
			responses.Unauthorized(response)

			return
		}
		if ok := authActions.ValidateJwt(authEntity.Hash); !ok {
			authActions.Logout(authEntity.UserId)
			responses.Unauthorized(response)

			return
		}

		ctx := context.WithValue(request.Context(), "user_id", authEntity.UserId)
		next.ServeHTTP(response, request.WithContext(ctx))
	})
}

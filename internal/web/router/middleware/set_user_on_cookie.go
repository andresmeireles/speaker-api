package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andresmeireles/speaker/internal/auth"
	"github.com/andresmeireles/speaker/internal/tools/responses"
)

func SetUserIdOnRequest(next http.Handler, repository auth.Repository) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("run set user on id migration")
		token, err := GetTokenFromRequest(req)
		if err != nil {
			responses.Unauthorized(res, err)

			return
		}

		auth, err := repository.GetByHash(token)
		if err != nil {
			responses.Unauthorized(res, err)

			return
		}

		ctx := context.WithValue(req.Context(), "user_id", auth.UserId)
		fmt.Println(ctx, ctx.Value("user_id"))
		next.ServeHTTP(res, req.WithContext(ctx))
	})
}

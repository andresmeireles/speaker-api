package router

import (
	"context"
	"net/http"

	"github.com/andresmeireles/speaker/internal/modules"
	"github.com/andresmeireles/speaker/internal/modules/auth"
	"github.com/andresmeireles/speaker/internal/modules/config"
	"github.com/andresmeireles/speaker/internal/modules/invite"
	"github.com/andresmeireles/speaker/internal/modules/person"
	"github.com/andresmeireles/speaker/internal/modules/user"
	"github.com/andresmeireles/speaker/internal/web/router/middleware"
	"github.com/go-chi/chi/v5"
)

var protectedModules = []modules.ModuleSetup{
	invite.Setup{},
	person.Setup{},
	config.Setup{},
	user.Setup{},
}

var openModules = []modules.ModuleSetup{
	auth.Setup{},
}

func Run(port string) {
	ctx := context.Background()
	router := chi.NewRouter()

	router.Use(middleware.ErrorHandler)
	router.Use(middleware.Cors)

	routes(ctx, router)

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		panic(err)
	}
}

func routes(ctx context.Context, router *chi.Mux) {
	router.Group(func(r chi.Router) {
		r.Use(middleware.CheckTokenOnCookie)
		for _, mod := range protectedModules {
			mod.Routes(r)
		}
		r.Get("/logout", func(w http.ResponseWriter, r *http.Request) {
			authController := auth.NewController()
			authController.Logout(w, r)
		})
	})

	for _, unprotectedRoutes := range openModules {
		unprotectedRoutes.Routes(router)
	}
}

package router

import (
	"context"
	"net/http"

	"github.com/andresmeireles/speaker/internal/auth"
	"github.com/andresmeireles/speaker/internal/config"
	"github.com/andresmeireles/speaker/internal/invite"
	"github.com/andresmeireles/speaker/internal/person"
	"github.com/andresmeireles/speaker/internal/stats"
	"github.com/andresmeireles/speaker/internal/tools/servicelocator"
	"github.com/andresmeireles/speaker/internal/user"
	"github.com/andresmeireles/speaker/internal/web"
	"github.com/andresmeireles/speaker/internal/web/router/middleware"
	"github.com/go-chi/chi/v5"
)

func Run(port string, services servicelocator.ServiceLocator) {
	ctx := context.Background()
	router := chi.NewRouter()

	router.Use(middleware.ErrorHandler)
	router.Use(middleware.Cors)

	routes(ctx, router, services)

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		panic(err)
	}
}

func routes(ctx context.Context, router *chi.Mux, sl servicelocator.ServiceLocator) {
	authController := servicelocator.Get[auth.AuthController](sl)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	router.Post("/login/confirm", authController.ReceiveCode)
	router.Post("/login", authController.ReceiveEmail)
	router.Post("/devtoken", authController.DevAuth)

	router.Group(func(r chi.Router) {
		r.Use(func(handler http.Handler) http.Handler {

		authActions := servicelocator.Get[auth.Actions](sl)
			return middleware.CheckTokenOnCookie(handler, authActions)
		})
		protectedRoutes(r, sl)
	})
}

//nolint:funlen
func protectedRoutes(r chi.Router, sl servicelocator.ServiceLocator) {
	authController := servicelocator.Get[auth.AuthController](sl)
	configController := servicelocator.Get[config.ConfigController](sl)
	speakerController := servicelocator.Get[person.PersonController](sl)
	userController := servicelocator.Get[user.UserController](sl)
	inviteController := servicelocator.Get[invite.InviteController](sl)
	statsController := servicelocator.Get[stats.StatsController](sl)

	r.Get("/logout", authController.Logout)

	r.Get("/configs", configController.GetConfigs)
	r.Post("/configs", configController.WriteConfig)

	r.Get("/speakers", speakerController.GetPersons)
	r.Get("/speakers/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := web.GetIntParameter(r, "id")
		speakerController.GetPerson(w, r, id)
	})
	r.Post("/speakers", speakerController.WritePerson)
	r.Put("/speakers", speakerController.UpdatePerson)
	r.Delete("/speakers", speakerController.DeletePerson)

	r.Get("/users/me", userController.Me)

	// INVITES
	r.Get("/invites", inviteController.GetAllInvites)
	r.Get("/invites/message/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := web.GetIntParameter(r, "id")
		inviteController.SendInvite(id, w, r)
	})
	r.Get("/invites/speaker/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := web.GetIntParameter(r, "id")
		inviteController.GetAllInvitesByPerson(w, r, id)
	})
	r.Get("/invites/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := web.GetIntParameter(r, "id")
		inviteController.GetInvite(id, w, r)
	})
	r.Post("/invites", inviteController.Create)
	// TODO: unificar alteracoes de status em uma rota
	r.Put("/invites/accept/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := web.GetIntParameter(r, "id")
		inviteController.Accepted(id, w, r)
	})
	r.Put("/invites/remember/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := web.GetIntParameter(r, "id")
		inviteController.Remember(id, w, r)
	})
	r.Put("/invites/done/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := web.GetIntParameter(r, "id")
		inviteController.WasDone(id, w, r)
	})
	r.Put("/invites/reject/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := web.GetIntParameter(r, "id")
		inviteController.Reject(w, r, id)
	})
	r.Put("/invites/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := web.GetIntParameter(r, "id")
		inviteController.Update(id, w, r)
	})
	r.Delete("/invites/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := web.GetIntParameter(r, "id")
		inviteController.DeleteInvite(w, r, id)
	})

	// statistics
	r.Get("/stats", statsController.SpeakersStats)
}

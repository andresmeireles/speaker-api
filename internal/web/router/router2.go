package router

import (
	"net/http"

	"github.com/andresmeireles/speaker/internal/auth"
	"github.com/andresmeireles/speaker/internal/config"
	"github.com/andresmeireles/speaker/internal/invite"
	"github.com/andresmeireles/speaker/internal/person"
	"github.com/andresmeireles/speaker/internal/stats"
	"github.com/andresmeireles/speaker/internal/user"
	"github.com/andresmeireles/speaker/internal/web"
	"github.com/andresmeireles/speaker/internal/web/router/middleware"
	"github.com/go-chi/chi/v5"
)

type Router struct {
	server            *chi.Mux
	authController    auth.AuthController
	configController  config.ConfigController
	speakerController person.PersonController
	userController    user.UserController
	inviteController  invite.InviteController
	statsController   stats.StatsController
	authActions       auth.Service
}

func NewRouter(
	server *chi.Mux,
	authController auth.AuthController,
	configController config.ConfigController,
	speakerController person.PersonController,
	userController user.UserController,
	inviteController invite.InviteController,
	statsController stats.StatsController,
	authActions auth.Actions,
) Router {
	return Router{
		server,
		authController,
		configController,
		speakerController,
		userController,
		inviteController,
		statsController,
		authActions,
	}
}

func (r Router) Run(port string) {
	r.server.Use(middleware.ErrorHandler)
	r.server.Use(middleware.Cors)

	r.nonAuthRoutes()
	r.authRoutes()

	if err := http.ListenAndServe(":"+port, r.server); err != nil {
		panic(err)
	}
}

func (r Router) nonAuthRoutes() {
	r.server.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	r.server.Post("/login/confirm", r.authController.ReceiveCode)
	r.server.Post("/login", r.authController.ReceiveEmail)
	r.server.Post("/devtoken", r.authController.DevAuth)
}

//nolint:funlen
func (r Router) authRoutes() {
	r.server.Group(func(router chi.Router) {
		router.Use(func(handler http.Handler) http.Handler {
			return middleware.CheckTokenOnCookie(handler, r.authActions)
		})

		router.Get("/logout", r.authController.Logout)

		router.Get("/configs", r.configController.GetConfigs)
		router.Post("/configs", r.configController.WriteConfig)

		router.Get("/speakers", r.speakerController.GetPersons)
		router.Get("/speakers/{id}", func(w http.ResponseWriter, req *http.Request) {
			id := web.GetIntParameter(req, "id")
			r.speakerController.GetPerson(w, req, id)
		})
		router.Post("/speakers", r.speakerController.WritePerson)
		router.Put("/speakers", r.speakerController.UpdatePerson)
		router.Delete("/speakers", r.speakerController.DeletePerson)

		router.Get("/users/me", r.userController.Me)

		// INVITES
		router.Get("/invites", r.inviteController.GetAllInvites)
		router.Get("/invites/message/{id}", func(w http.ResponseWriter, req *http.Request) {
			id := web.GetIntParameter(req, "id")
			r.inviteController.SendInvite(id, w, req)
		})
		router.Get("/invites/speaker/{id}", func(w http.ResponseWriter, req *http.Request) {
			id := web.GetIntParameter(req, "id")
			r.inviteController.GetAllInvitesByPerson(w, req, id)
		})
		router.Get("/invites/{id}", func(w http.ResponseWriter, req *http.Request) {
			id := web.GetIntParameter(req, "id")
			r.inviteController.GetInvite(id, w, req)
		})
		router.Post("/invites", r.inviteController.Create)
		// TODO: unificar alteracoes de status em uma rota
		router.Put("/invites/accept/{id}", func(w http.ResponseWriter, req *http.Request) {
			id := web.GetIntParameter(req, "id")
			r.inviteController.Accepted(id, w, req)
		})
		router.Put("/invites/remember/{id}", func(w http.ResponseWriter, req *http.Request) {
			id := web.GetIntParameter(req, "id")
			r.inviteController.Remember(id, w, req)
		})
		router.Put("/invites/done/{id}", func(w http.ResponseWriter, req *http.Request) {
			id := web.GetIntParameter(req, "id")
			r.inviteController.WasDone(id, w, req)
		})
		router.Put("/invites/reject/{id}", func(w http.ResponseWriter, req *http.Request) {
			id := web.GetIntParameter(req, "id")
			r.inviteController.Reject(w, req, id)
		})
		router.Put("/invites/{id}", func(w http.ResponseWriter, req *http.Request) {
			id := web.GetIntParameter(req, "id")
			r.inviteController.Update(id, w, req)
		})
		router.Delete("/invites/{id}", func(w http.ResponseWriter, req *http.Request) {
			id := web.GetIntParameter(req, "id")
			r.inviteController.DeleteInvite(w, req, id)
		})

		// statistics
		router.Get("/stats", r.statsController.SpeakersStats)
	})
}

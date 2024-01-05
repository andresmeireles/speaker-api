package router

import (
	"net/http"
	"strconv"

	"github.com/andresmeireles/speaker/internal/modules/config"
	"github.com/andresmeireles/speaker/internal/modules/invite"
	"github.com/andresmeireles/speaker/internal/modules/person"
	"github.com/andresmeireles/speaker/internal/router/middleware"
	"github.com/go-chi/chi/v5"
)

func Run(port string) {
	router := chi.NewRouter()

	router.Use(middleware.Cors)

	routes(router)

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		panic(err)
	}
}

func guardRoutes(router chi.Router) {
	router.Get("/invites", invite.GetAllInvites)
	router.Get("/invites/message/{id}", invite.SendInvite)
	router.Post("/invites", invite.Create)
	router.Put("/invites/accept/{id}", invite.Accepted)
	router.Put("/invites/remember/{id}", invite.Remember)
	router.Put("/invites", func(w http.ResponseWriter, r *http.Request) {
		inviteIdParam := chi.URLParam(r, "inviteId")
		inviteId, err := strconv.Atoi(inviteIdParam)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("bad formatted url"))

			return
		}

		invite.Update(inviteId, w, r)
	})
	router.Delete("/invites/{id}", invite.DeleteInvite)

	router.Get("/configs", config.GetConfigs)
	router.Post("/configs", config.WriteConfig)

	router.Get("/speakers", person.GetPersons)
	router.Post("/speakers", person.WritePerson)
	router.Put("/speakers", person.UpdatePerson)
	router.Delete("/speakers", person.DeletePerson)
}

func routes(router *chi.Mux) {
	// router.Get("/", person.ShowMode)

	router.Group(func(r chi.Router) {
		// r.Use(middleware.CheckTokenOnCookie)
		guardRoutes(r)
	})
}

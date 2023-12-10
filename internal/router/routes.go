package router

import (
	"net/http"
	"strconv"

	"github.com/andresmeireles/speaker/internal/modules/invite"
	"github.com/andresmeireles/speaker/internal/modules/person"
	"github.com/andresmeireles/speaker/internal/router/middleware"
	"github.com/go-chi/chi/v5"
)

func Run(port string) {
	router := chi.NewRouter()

	routes(router)

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		panic(err)
	}
}

func guardRoutes(router chi.Router) {
	router.Get("/invites", invite.GetAllInvites)
	router.Post("/invite", invite.Create)
	router.Put("/invite", func(w http.ResponseWriter, r *http.Request) {
		inviteIdParam := chi.URLParam(r, "inviteId")
		inviteId, err := strconv.Atoi(inviteIdParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("bad formatted url"))
			return
		}
		invite.Update(inviteId, w, r)
	})
}

func routes(router *chi.Mux) {
	router.Get("/", person.ShowMode)

	router.Group(func(r chi.Router) {
		r.Use(middleware.CheckTokenOnCookie)
		guardRoutes(r)
	})
}

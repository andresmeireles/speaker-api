package invite

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/andresmeireles/speaker/internal/modules/config"
	"github.com/andresmeireles/speaker/internal/modules/person"
	"github.com/go-chi/chi/v5"
)

type Setup struct{}

func (s Setup) Routes(router chi.Router) {
	controller := NewController()

	router.Get("/invites", controller.GetAllInvites)
	router.Get("/invites/message/{id}", func(w http.ResponseWriter, r *http.Request) {
		inviteId := chi.URLParam(r, "id")
		id, err := strconv.Atoi(inviteId)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("bad formatted url"))

			return
		}

		controller.SendInvite(id, w, r)
	})
	router.Post("/invites", controller.Create)
	router.Put("/invites/accept/{id}", func(w http.ResponseWriter, r *http.Request) {
		inviteIdParam := chi.URLParam(r, "id")
		inviteId, err := strconv.Atoi(inviteIdParam)

		if err != nil {
			slog.Error("error on accepted invite controller, error on decode", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("bad formatted url"))

			return
		}

		controller.Accepted(inviteId, w, r)
	})
	router.Put("/invites/remember/{id}", func(w http.ResponseWriter, r *http.Request) {
		inviteIdParam := chi.URLParam(r, "id")
		inviteId, err := strconv.Atoi(inviteIdParam)

		if err != nil {
			slog.Error("error on accepted invite controller, error on decode", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("bad formatted url"))

			return
		}

		controller.Remember(inviteId, w, r)
	})
	router.Put("/invites", func(w http.ResponseWriter, r *http.Request) {
		inviteIdParam := chi.URLParam(r, "inviteId")
		inviteId, err := strconv.Atoi(inviteIdParam)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("bad formatted url"))

			return
		}

		controller.Update(inviteId, w, r)
	})
	router.Delete("/invites/{id}", func(w http.ResponseWriter, r *http.Request) {
		inviteIdParam := chi.URLParam(r, "inviteId")
		inviteId, err := strconv.Atoi(inviteIdParam)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("bad formatted url"))

			return
		}

		controller.DeleteInvite(inviteId, w, r)
	})
}

func NewController() InviteController {
	actions := NewActions()

	return InviteController{
		inviteRepository: InviteRepository{},
		personRepository: person.PersonRepository{},
		configRepository: config.ConfigRepository{},
		action:           actions,
	}
}

func NewActions() Actions {
	return Actions{
		inviteRepository: InviteRepository{},
		personRepository: person.PersonRepository{},
		configRepository: config.ConfigRepository{},
	}
}

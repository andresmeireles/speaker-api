package invite

import (
	"net/http"

	"github.com/andresmeireles/speaker/internal/modules/config"
	"github.com/andresmeireles/speaker/internal/modules/person"
	"github.com/andresmeireles/speaker/internal/web"
	"github.com/go-chi/chi/v5"
)

type Setup struct{}

func (s Setup) Routes(router chi.Router) {
	controller := NewController()
	getRoutes(router, controller)
	postRoutes(router, controller)
	putRoutes(router, controller)
	deleteRoutes(router, controller)
}

func getRoutes(router chi.Router, controller InviteController) {
	router.Get("/invites", controller.GetAllInvites)
	router.Get("/invites/message/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err, handlerFunc := web.GetIntParameter(r, "id")
		if err != nil {
			handlerFunc(w)
		} else {
			controller.SendInvite(id, w, r)
		}
	})
}

func postRoutes(router chi.Router, controller InviteController) {
	router.Post("/invites", controller.Create)
}

func putRoutes(router chi.Router, controller InviteController) {
	router.Put("/invites/accept/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err, handlerFunc := web.GetIntParameter(r, "id")
		if err != nil {
			handlerFunc(w)
		} else {
			controller.Accepted(id, w, r)
		}
	})
	router.Put("/invites/remember/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err, handlerFunc := web.GetIntParameter(r, "id")
		if err != nil {
			handlerFunc(w)
		} else {
			controller.Remember(id, w, r)
		}
	})
	router.Put("/invites", func(w http.ResponseWriter, r *http.Request) {
		id, err, handlerFunc := web.GetIntParameter(r, "id")
		if err != nil {
			handlerFunc(w)
		} else {
			controller.Update(id, w, r)
		}
	})
}

func deleteRoutes(router chi.Router, controller InviteController) {
	router.Delete("/invites/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err, handlerFunc := web.GetIntParameter(r, "id")
		if err != nil {
			handlerFunc(w)
		} else {
			controller.DeleteInvite(id, w, r)
		}
	})
}

func NewController() InviteController {
	return InviteController{
		inviteRepository: NewRepository(),
		personRepository: person.PersonRepository{},
		configRepository: config.ConfigRepository{},
		action:           NewActions(),
	}
}

func NewRepository() InviteRepository {
	return InviteRepository{
		personRepository: person.PersonRepository{},
	}
}

func NewActions() Actions {
	return Actions{
		inviteRepository: NewRepository(),
		personRepository: person.PersonRepository{},
		configRepository: config.ConfigRepository{},
	}
}

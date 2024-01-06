package auth

import "github.com/go-chi/chi/v5"

type Setup struct{}

func (s Setup) Routes(router chi.Router) {
	controller := NewController()

	router.Get("/login", controller.ReceiveEmail)
}

func NewController() AuthController {
	return AuthController{}
}

func NewActions() Actions {
	return Actions{}
}

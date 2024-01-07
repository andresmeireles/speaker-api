package user

import "github.com/go-chi/chi/v5"

type Setup struct{}

func (s Setup) Routes(router chi.Router) {
	controller := NewController()
	router.Get("/users/me", controller.Me)
}

func NewController() UserController {
	return UserController{
		repository: UserRepository{},
	}
}

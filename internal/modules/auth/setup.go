package auth

import (
	"github.com/andresmeireles/speaker/internal/modules/codesender"
	"github.com/andresmeireles/speaker/internal/modules/user"
	"github.com/andresmeireles/speaker/internal/tools"
	"github.com/go-chi/chi/v5"
)

type Setup struct{}

func (s Setup) Routes(router chi.Router) {
	controller := NewController()

	router.Post("/login", controller.ReceiveEmail)
	router.Post("/confirm", controller.ReceiveCode)
}

func NewController() AuthController {
	return AuthController{
		actions: NewActions(),
	}
}

func NewActions() Actions {
	email, err := tools.NewDefaultEmail()
	if err != nil {
		panic(err)
	}

	return Actions{
		repository:       AuthRepository{},
		userRepository:   user.UserRepository{},
		email:            email,
		codeSenderAction: codesender.NewActions(),
	}
}

package config

import "github.com/go-chi/chi/v5"

type Setup struct{}

func (s Setup) Routes(router chi.Router) {
	controller := NewController()
	router.Get("/configs", controller.GetConfigs)
	router.Post("/configs", controller.WriteConfig)
}

func NewController() ConfigController {
	return ConfigController{}
}

func NewActions() Actions {
	return Actions{
		configRepository: ConfigRepository{},
	}
}

package person

import "github.com/go-chi/chi/v5"

type Setup struct{}

func (s Setup) Routes(router chi.Router) {
	controller := NewController()

	router.Get("/speakers", controller.GetPersons)
	router.Post("/speakers", controller.WritePerson)
	router.Put("/speakers", controller.UpdatePerson)
	router.Delete("/speakers", controller.DeletePerson)
}

func NewAction() Actions {
	return Actions{
		repository: PersonRepository{},
	}
}

func NewController() PersonController {
	return PersonController{
		personRepository: PersonRepository{},
		actions:          NewAction(),
	}
}

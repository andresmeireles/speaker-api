package codesender

import "github.com/go-chi/chi/v5"

type Setup struct{}

func (s Setup) Routes(router chi.Router) {}

func NewActions() Actions {
	return Actions{
		repository: AuthCodeRepository{},
	}
}

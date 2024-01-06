package modules

import "github.com/go-chi/chi/v5"

type ModuleSetup interface {
	Routes(router chi.Router)
}

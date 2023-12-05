package router

import (
	"net/http"

	"github.com/andresmeireles/speaker/internal/modules/person"
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

func routes(router *chi.Mux) {
	router.Get("/", person.ShowMode)
}

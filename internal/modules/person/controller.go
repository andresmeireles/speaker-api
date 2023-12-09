package person

import (
	"net/http"
	"os"

	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/web"
)

func ShowMode(w http.ResponseWriter, r *http.Request) {
	mode := os.Getenv("MODE")

	w.Write([]byte(mode))
}

func AddPerson(w http.ResponseWriter, r *http.Request) {
	person, err := web.DecodePostBody[entity.Person](r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	_, err = AddNewPerson(person)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Person created"))
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {}

func DeletePerson(w http.ResponseWriter, r *http.Request) {}

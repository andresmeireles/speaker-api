package person

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/logger"
	"github.com/andresmeireles/speaker/internal/web"
)

func ShowMode(w http.ResponseWriter, r *http.Request) {
	mode := os.Getenv("MODE")
	logger.Info("super", "mode")

	w.Write([]byte(mode))
}

func GetPersons(w http.ResponseWriter, r *http.Request) {
	repo := PersonRepository{}
	persons, err := repo.GetAll()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	response, err := json.Marshal(persons)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func WritePerson(w http.ResponseWriter, r *http.Request) {
	person, err := web.DecodePostBody[entity.Person](r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	fmt.Println(person)

	err = Write(person, PersonRepository{})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Person created"))
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	personId, err := web.DecodePostBody[DeletePersonData](r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error on decode"))
		return
	}

	id, err := strconv.Atoi(personId.Speaker)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error on decode"))
		return
	}

	person, err := PersonRepository{}.GetById(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error on decode"))
		return
	}

	err = PersonRepository{}.Delete(*person)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error on decode"))
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("removed person"))
}

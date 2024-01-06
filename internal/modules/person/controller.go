package person

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/logger"
	web "github.com/andresmeireles/speaker/internal/web/decoder"
)

type PersonController struct {
	personRepository PersonRepository
	actions          Actions
}

func (p PersonController) GetPersons(w http.ResponseWriter, r *http.Request) {
	persons, err := p.personRepository.GetAll()
	if err != nil {
		slog.Error("error on get persons", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		return
	}

	response, err := json.Marshal(persons)
	if err != nil {
		slog.Error("error on get persons", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (p PersonController) WritePerson(w http.ResponseWriter, r *http.Request) {
	person, err := web.DecodePostBody[entity.Person](r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	err = p.actions.Write(person)
	if err != nil {
		slog.Error("error on write person", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Person created"))
}

func (p PersonController) DeletePerson(w http.ResponseWriter, r *http.Request) {
	personId, err := web.DecodePostBody[DeletePersonData](r.Body)
	if err != nil {
		slog.Error("error on decode", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error on decode"))

		return
	}

	id, err := strconv.Atoi(personId.Speaker)
	if err != nil {
		slog.Error("error on decode", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error on decode"))

		return
	}

	repository := p.personRepository
	person, err := repository.GetById(id)

	if err != nil {
		slog.Error("error on decode", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error on decode"))

		return
	}

	err = repository.Delete(*person)
	if err != nil {
		slog.Error("error on decode", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error on decode"))

		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("removed person"))
}

func (p PersonController) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	person, err := web.DecodePostBody[entity.Person](r.Body)
	if err != nil {
		logger.Error("error cannot decode", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	personRepo := p.personRepository
	if err = personRepo.Update(person); err != nil {
		slog.Error("error cannot update", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Person successfully updated"))
}

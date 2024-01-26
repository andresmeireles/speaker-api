package person

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/andresmeireles/speaker/internal/tools/responses"
	web "github.com/andresmeireles/speaker/internal/web/decoder"
)

type PersonController struct {
	personRepository PersonRepository
	actions          ActionsInterface
}

func NewController(repository PersonRepository, actions ActionsInterface) PersonController {
	return PersonController{
		personRepository: repository,
		actions:          actions,
	}
}

func (c PersonController) GetPerson(w http.ResponseWriter, r *http.Request, personId int) {
	person, err := c.personRepository.GetById(personId)
	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	response, err := json.Marshal(person)
	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	responses.Ok(w, response)
}

func (p PersonController) GetPersons(w http.ResponseWriter, r *http.Request) {
	persons, err := p.personRepository.GetAll()
	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	response, err := json.Marshal(persons)
	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	responses.Ok(w, response)
}

func (p PersonController) WritePerson(w http.ResponseWriter, r *http.Request) {
	person, err := web.DecodePostBody[Person](r.Body)
	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	err = p.actions.Write(person)
	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	responses.Created(w, []byte("Person successfully created"))
}

func (p PersonController) DeletePerson(w http.ResponseWriter, r *http.Request) {
	personId, err := web.DecodePostBody[DeletePersonData](r.Body)
	if err != nil {
		responses.DecodeError(w, err)

		return
	}

	id, err := strconv.Atoi(personId.Speaker)
	if err != nil {
		responses.DecodeError(w, err)

		return
	}

	repository := p.personRepository
	person, err := repository.GetById(id)

	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	err = repository.Delete(*person)
	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	responses.Ok(w, []byte("Person successfully deleted"))
}

func (p PersonController) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	person, err := web.DecodePostBody[Person](r.Body)
	if err != nil {
		responses.DecodeError(w, err)

		return
	}

	personRepo := p.personRepository
	if err = personRepo.Update(person); err != nil {
		responses.BadResponse(w, err)

		return
	}

	responses.Ok(w, []byte("Person successfully updated"))
}

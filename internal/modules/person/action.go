package person

import (
	"database/sql"

	"github.com/andresmeireles/speaker/internal/db/entity"
)

type Actions struct {
	repository PersonRepository
}

func (a Actions) Write(person entity.Person) error {
	dbPerson, err := a.repository.GetByName(person.Name)
	if err == sql.ErrNoRows {
		return a.repository.Add(person)
	}

	if err != nil {
		return err
	}

	return a.repository.Update(*dbPerson)
}

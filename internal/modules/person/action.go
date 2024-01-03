package person

import (
	"database/sql"

	"github.com/andresmeireles/speaker/internal/db/entity"
)

func Write(person entity.Person, repository PersonRepository) error {
	dbPerson, err := repository.GetByName(person.Name)

	if err == sql.ErrNoRows {
		return repository.Add(person)
	}

	if err != nil {
		return err
	}

	return repository.Update(*dbPerson)
}

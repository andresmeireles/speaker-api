package person

import (
	"database/sql"

	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/tools/servicelocator"
)

type Actions struct {
	repository PersonRepository
}

func (a Actions) New(s servicelocator.ServiceLocator) any {
	return Actions{
		repository: servicelocator.Get[PersonRepository](s),
	}
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

func (a Actions) RemovePerson(personId int) error {
	return a.repository.Delete(entity.Person{Id: personId})
}

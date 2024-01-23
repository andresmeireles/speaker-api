package person

import (
	"database/sql"
)

type ActionsInterface interface {
	Write(person Person) error
	RemovePerson(personId int) error
}

type Actions struct {
	repository Repository
}

func NewAction(repository Repository) Actions {
	return Actions{
		repository: repository,
	}
}

func (a Actions) Write(person Person) error {
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
	return a.repository.Delete(Person{Id: personId})
}

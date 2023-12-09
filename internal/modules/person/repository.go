package person

import (
	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/db/repository"
)

type PersonRepository struct{}

func (r PersonRepository) Add(person entity.Person) error {
	return repository.Add(person)
}

func (r PersonRepository) GetById(id int) (*entity.Person, error) {
	return repository.GetById[entity.Person](id, entity.Person{})
}

func (r PersonRepository) Update(person entity.Person) error {
	return repository.Update(person)
}

func (r PersonRepository) GetAll() []entity.Person {
	return repository.GetAll[entity.Person](entity.Person{})
}

func (r PersonRepository) Delete(person entity.Person) error {
	return repository.Delete(person)
}

func AddNewPerson(person entity.Person) (bool, error) {
	err := repository.Add(person)
	if err != nil {
		return false, err
	}
	return true, nil
}

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
	row := repository.GetById[entity.Person](id)
	person := new(entity.Person)

	if err := row.Scan(
		&person.Id,
		&person.Name,
	); err != nil {
		return nil, err
	}

	return person, nil

}

func (r PersonRepository) Update(person entity.Person) error {
	return repository.Update(person)
}

func (r PersonRepository) GetAll() ([]entity.Person, error) {
	rows, err := repository.GetAll[entity.Person]()

	if err != nil {
		return nil, err
	}

	people := make([]entity.Person, 0)

	for rows.Next() {
		person := new(entity.Person)

		if err := rows.Scan(
			&person.Id,
			&person.Name,
		); err != nil {
			return nil, err
		}

		people = append(people, *person)
	}

	return people, nil
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

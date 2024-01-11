package person

import (
	"fmt"

	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/db/repository"
	"github.com/andresmeireles/speaker/internal/tools/servicelocator"
)

type PersonRepository struct {
	repository repository.Repository[entity.Person]
}

func (r PersonRepository) New(s servicelocator.ServiceLocator) any {
	return PersonRepository{
		repository: servicelocator.Get[repository.Repository[entity.Person]](s),
	}
}

func (r PersonRepository) Add(person entity.Person) error {
	return r.repository.Add(person)
}

func (r PersonRepository) GetById(id int) (*entity.Person, error) {
	row, err := r.repository.GetById(id)
	if err != nil {
		return nil, err
	}

	person := new(entity.Person)
	if err := row.Scan(
		&person.Id,
		&person.Name,
	); err != nil {
		return nil, err
	}

	return person, nil
}

func (r PersonRepository) GetByName(name string) (*entity.Person, error) {
	person := new(entity.Person)
	query := fmt.Sprintf("SELECT * FROM %s WHERE name = $1 LIMIT 1", person.Table())
	row, err := r.repository.SingleQuery(query, name)

	if err != nil {
		return nil, err
	}

	if err := row.Scan(&person.Id, &person.Name); err != nil {
		return nil, err
	}

	return person, nil
}

func (r PersonRepository) Update(person entity.Person) error {
	return r.repository.Update(person)
}

func (r PersonRepository) GetAll() ([]entity.Person, error) {
	rows, err := r.repository.GetAll()
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
	return r.repository.Delete(person)
}

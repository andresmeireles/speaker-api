package person

import (
	"fmt"

	"github.com/andresmeireles/speaker/internal/repository"
	"github.com/andresmeireles/speaker/internal/tools/servicelocator"
)

type PersonRepository struct {
	repository repository.Repository
}

type PersonRepositoryInterface interface {
	Add(person Person) error
	GetById(id int) (*Person, error)
	GetByName(name string) (*Person, error)
	Update(person Person) error
	GetAll() ([]Person, error)
	Delete(person Person) error
}

func (r PersonRepository) New(s servicelocator.ServiceLocator) any {
	return PersonRepository{
		repository: servicelocator.Get[repository.Repository](s),
	}
}

func (r PersonRepository) Add(person Person) error {
	return r.repository.Add(person)
}

func (r PersonRepository) GetById(id int) (*Person, error) {
	person := new(Person)
	row, err := r.repository.GetById(person.Table(), id)

	if err != nil {
		return nil, err
	}

	if err := row.Scan(
		&person.Id,
		&person.Name,
	); err != nil {
		return nil, err
	}

	return person, nil
}

func (r PersonRepository) GetByName(name string) (*Person, error) {
	person := new(Person)
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

func (r PersonRepository) Update(person Person) error {
	return r.repository.Update(person)
}

func (r PersonRepository) GetAll() ([]Person, error) {
	rows, err := r.repository.GetAll(Person{}.Table())
	if err != nil {
		return nil, err
	}

	people := make([]Person, 0)

	for rows.Next() {
		person := new(Person)
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

func (r PersonRepository) Delete(person Person) error {
	return r.repository.Delete(person)
}

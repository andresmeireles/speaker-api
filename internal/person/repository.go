package person

import (
	"fmt"

	"github.com/andresmeireles/speaker/internal/repository"
	"github.com/andresmeireles/speaker/internal/tools/servicelocator"
)

type PersonRepository interface {
	Add(person Person) error
	GetById(id int) (*Person, error)
	GetByName(name string) (*Person, error)
	Update(person Person) error
	GetAll() ([]Person, error)
	Delete(person Person) error
}

type Repository struct {
	repository repository.Repository
}

func NewRepository(repository repository.Repository) PersonRepository {
	return Repository{
		repository: repository,
	}
}

func (r Repository) New(s servicelocator.ServiceLocator) any {
	return Repository{
		repository: servicelocator.Get[repository.Repository](s),
	}
}

func (r Repository) Add(person Person) error {
	return r.repository.Add(person)
}

func (r Repository) GetById(id int) (*Person, error) {
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

func (r Repository) GetByName(name string) (*Person, error) {
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

func (r Repository) Update(person Person) error {
	return r.repository.Update(person)
}

func (r Repository) GetAll() ([]Person, error) {
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

func (r Repository) Delete(person Person) error {
	return r.repository.Delete(person)
}

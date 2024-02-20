package person

import (
	"fmt"

	"github.com/andresmeireles/speaker/internal/db"
	"github.com/andresmeireles/speaker/internal/repository"
)

type PersonRepository interface {
	Add(person Person) error
	GetById(id int) (*Person, error)
	GetByName(name string) (*Person, error)
	Update(person Person) error
	GetAll() ([]Person, error)
	Delete(person Person) error
}

const tableName = "persons"

type Repository struct {
	repository repository.Repository
}

func NewRepository(repository repository.Repository) PersonRepository {
	return Repository{
		repository: repository,
	}
}

func (r Repository) Add(person Person) error {
	return r.repository.Add(person)
}

func (r Repository) GetById(id int) (*Person, error) {
	row, err := r.repository.GetById(tableName, id)

	if err != nil {
		return nil, err
	}

	person, err := r.scan(row)
	if err != nil {
		return nil, err
	}

	return person, nil
}

func (r Repository) GetByName(name string) (*Person, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE name = $1 LIMIT 1", tableName)
	row, err := r.repository.SingleQuery(query, name)

	if err != nil {
		return nil, err
	}

	person, err := r.scan(row)
	if err != nil {
		return nil, err
	}

	return person, nil
}

func (r Repository) GetAll() ([]Person, error) {
	rows, err := r.repository.GetAll(tableName)
	if err != nil {
		return nil, err
	}

	people := make([]Person, 0)

	for rows.Next() {
		person, err := r.scan(rows)
		if err != nil {
			return nil, err
		}

		people = append(people, *person)
	}

	return people, nil
}

func (r Repository) scan(row db.RowScanner) (*Person, error) {
	person := new(Person)
	if err := row.Scan(&person.Id, &person.Name, &person.LastName, &person.Gender); err != nil {
		return nil, err
	}

	return person, nil
}

func (r Repository) Update(person Person) error {
	return r.repository.Update(person)
}

func (r Repository) Delete(person Person) error {
	return r.repository.Delete(person)
}

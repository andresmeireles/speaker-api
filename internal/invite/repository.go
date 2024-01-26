package invite

import (
	"database/sql"
	"fmt"
	"log/slog"
	"sort"

	"github.com/andresmeireles/speaker/internal/person"
	"github.com/andresmeireles/speaker/internal/repository"
)

type InviteRepository interface {
	Add(invite Invite) error
	Query(query string, values ...any) (*sql.Rows, error)
	GetAllOrdered(field string, asc bool) ([]Invite, error)
	GetByPersonId(id int) ([]Invite, error)
	GetById(id int) (*Invite, error)
	Update(invite Invite) error
	UpdateStatus(invite Invite, status int) error
	Delete(invite Invite) error
}

type Repository struct {
	repository       repository.Repository
	personRepository person.PersonRepository
}

func NewRepository(repository repository.Repository) Repository {
	return Repository{
		repository:       repository,
		personRepository: person.NewRepository(repository),
	}
}

func (r Repository) Add(invite Invite) error {
	return r.repository.Add(invite)
}

func (r Repository) Query(query string, values ...any) (*sql.Rows, error) {
	return r.repository.Query(query, values...)
}

func (r Repository) GetAllOrdered(field string, asc bool) ([]Invite, error) {
	invites := make([]Invite, 0)
	query := "SELECT * FROM invites ORDER BY "

	if asc {
		query += field + " ASC"
	} else {
		query += field + " DESC"
	}

	rows, err := r.repository.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		invite, err := r.scan(rows.Scan)
		if err != nil {
			return nil, err
		}

		invites = append(invites, *invite)
	}

	return invites, nil
}

func (r Repository) GetAll() ([]Invite, error) {
	rows, err := r.repository.GetAll(Invite{}.Table())
	if err != nil {
		return nil, err
	}

	invites := make([]Invite, 0)

	for rows.Next() {
		invite, err := r.scan(rows.Scan)
		if err != nil {
			return nil, err
		}

		invites = append(invites, *invite)
	}

	return invites, nil
}

func (r Repository) GetById(id int) (*Invite, error) {
	query := "SELECT * FROM invites WHERE id = $1 ORDER BY date ASC"
	row, err := r.repository.Query(query, id)

	if err != nil {
		return nil, err
	}

	invite, err := r.scan(row.Scan)
	if err != nil {
		return nil, err
	}

	return invite, nil
}

func (r Repository) GetByPersonId(id int) ([]Invite, error) {
	invites := make([]Invite, 0)
	query := "SELECT * FROM invites WHERE person_id = $1"
	rows, err := r.repository.Query(query, id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		invite, err := r.scan(rows.Scan)
		if err != nil {
			return nil, err
		}

		invites = append(invites, *invite)
	}

	return invites, nil
}

func (r Repository) scan(scanFunc func(dest ...any) error) (*Invite, error) {
	invite := new(Invite)
	if err := scanFunc(
		&invite.Id,
		&invite.Theme,
		&invite.References,
		&invite.Date,
		&invite.Time,
		&invite.PersonId,
		&invite.Status,
	); err != nil {
		return nil, err
	}

	person, err := r.personRepository.GetById(invite.PersonId)
	if err != nil {
		return nil, err
	}

	invite.Person = *person

	return invite, nil
}

func (r Repository) Update(invite Invite) error {
	return r.repository.Update(invite)
}

func (r Repository) Delete(invite Invite) error {
	invitesPerPerson, err := r.GetByPersonId(invite.PersonId)
	if err != nil {
		return err
	}

	if len(invitesPerPerson) > 0 {
		slog.Error("error on delete invite, person has more than one invite")

		return fmt.Errorf("person has more than one invite")
	}

	return r.repository.Delete(invite)
}

func (r Repository) UpdateStatus(invite Invite, status int) error {
	statuses := Invite{}.StatusList()
	index := sort.SearchInts(statuses, status)
	isAllowedIndex := index < len(statuses) && statuses[index] == status

	if !isAllowedIndex {
		return fmt.Errorf("status not allowed")
	}

	invite.Status = status

	return r.Update(invite)
}

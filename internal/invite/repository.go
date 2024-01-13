package invite

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/andresmeireles/speaker/internal/person"
	"github.com/andresmeireles/speaker/internal/repository"
	"github.com/andresmeireles/speaker/internal/tools/servicelocator"
)

type InviteRepository struct {
	repository       repository.Repository
	personRepository person.PersonRepository
}

func NewRepository(repository repository.Repository) InviteRepository {
	return InviteRepository{
		repository:       repository,
		personRepository: person.NewRepository(repository),
	}
}

func (r InviteRepository) New(s servicelocator.ServiceLocator) any {
	return InviteRepository{
		repository:       servicelocator.Get[repository.Repository](s),
		personRepository: servicelocator.Get[person.PersonRepository](s),
	}
}

func (r InviteRepository) Add(invite Invite) error {
	return r.repository.Add(invite)
}

func (r InviteRepository) Query(query string, values ...any) (*sql.Rows, error) {
	return r.repository.Query(query, values...)
}

func (r InviteRepository) GetAllOrdered(field string, asc bool) ([]Invite, error) {
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

func (r InviteRepository) GetAll() ([]Invite, error) {
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

func (r InviteRepository) GetById(id int) (*Invite, error) {
	row, err := r.repository.GetById(Invite{}.Table(), id)
	if err != nil {
		return nil, err
	}

	invite, err := r.scan(row.Scan)
	if err != nil {
		return nil, err
	}

	return invite, nil
}

func (r InviteRepository) GetByPersonId(id int) ([]Invite, error) {
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

func (r InviteRepository) scan(scanFunc func(dest ...any) error) (*Invite, error) {
	invite := new(Invite)
	if err := scanFunc(
		&invite.Id,
		&invite.Theme,
		&invite.References,
		&invite.Date,
		&invite.Time,
		&invite.Accepted,
		&invite.Remembered,
		&invite.PersonId,
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

func (r InviteRepository) Update(invite Invite) error {
	return r.repository.Update(invite)
}

func (r InviteRepository) Delete(invite Invite) error {
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

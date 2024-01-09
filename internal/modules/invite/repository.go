package invite

import (
	"database/sql"
	"fmt"

	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/db/repository"
	"github.com/andresmeireles/speaker/internal/modules/person"
)

type InviteRepository struct {
	personRepository person.PersonRepository
}

func (r InviteRepository) Add(invite entity.Invite) error {
	return repository.Add(invite)
}

func (r InviteRepository) Query(query string, values ...any) (*sql.Rows, error) {
	return repository.Query(query, values...)
}

func (r InviteRepository) GetAllOrdered(field string, asc bool) ([]entity.Invite, error) {
	invites := make([]entity.Invite, 0)
	query := "SELECT * FROM invites ORDER BY "

	if asc {
		query += field + " ASC"
	} else {
		query += field + " DESC"
	}

	rows, err := repository.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		invite := new(entity.Invite)
		if err := rows.Scan(
			&invite.Id,
			&invite.Theme,
			&invite.References,
			&invite.Date,
			&invite.Time,
			&invite.Accepted,
			&invite.Remembered,
			&invite.PersonId,
		); err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("no rows found")
			}

			return nil, err
		}
		person, err := r.personRepository.GetById(invite.PersonId)
		if err != nil {
			return nil, err
		}
		invite.Person = *person

		invites = append(invites, *invite)
	}

	return invites, nil
}

func (r InviteRepository) GetAll() ([]entity.Invite, error) {
	invites := make([]entity.Invite, 0)
	rows, err := repository.GetAll[entity.Invite]()
	personRepository := person.PersonRepository{}

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		invite := new(entity.Invite)
		if err := rows.Scan(
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

		person, err := personRepository.GetById(invite.PersonId)
		if err != nil {
			return nil, err
		}

		invite.Person = *person
		invites = append(invites, *invite)
	}

	return invites, nil
}

func (r InviteRepository) GetById(id int) (*entity.Invite, error) {
	invite := new(entity.Invite)
	row, err := repository.GetById[entity.Invite](id)

	if err != nil {
		return nil, err
	}

	personRepository := person.PersonRepository{}

	if err := row.Scan(
		&invite.Id,
		&invite.Theme,
		&invite.References,
		&invite.Date,
		&invite.Time,
		&invite.Remembered,
		&invite.Accepted,
		&invite.PersonId,
	); err != nil {
		return nil, err
	}

	person, err := personRepository.GetById(invite.PersonId)
	if err != nil {
		return nil, err
	}

	invite.Person = *person

	return invite, nil
}

func (r InviteRepository) Update(invite entity.Invite) error {
	return repository.Update(invite)
}

func (r InviteRepository) Delete(invite entity.Invite) error {
	return repository.Delete(invite)
}

package invite

import (
	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/db/repository"
)

type InviteRepository struct{}

func (r InviteRepository) Add(invite entity.Invite) error {
	return repository.Add(invite)
}

func (r InviteRepository) GetAll() ([]entity.Invite, error) {
	invites := make([]entity.Invite, 0)
	rows, err := repository.GetAll[entity.Invite]()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		invite := new(entity.Invite)

		if err := rows.Scan(
			&invite.Id,
			&invite.Person,
			&invite.Theme,
			&invite.Date,
			&invite.Time,
			&invite.Remembered,
			&invite.Accepted,
		); err != nil {
			return nil, err
		}

		invites = append(invites, *invite)
	}

	return invites, nil
}

func (r InviteRepository) GetById(id int) (*entity.Invite, error) {
	invite := new(entity.Invite)
	row := repository.GetById[entity.Invite](id)

	if err := row.Scan(
		&invite.Id,
		&invite.Person,
		&invite.Theme,
		&invite.Date,
		&invite.Time,
		&invite.Remembered,
		&invite.Accepted,
	); err != nil {
		return nil, err
	}

	return invite, nil
}

func (r InviteRepository) Update(invite entity.Invite) error {
	return repository.Update(invite)
}

func (r InviteRepository) Delete(invite entity.Invite) error {
	return repository.Delete(invite)
}

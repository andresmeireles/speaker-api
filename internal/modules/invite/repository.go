package invite

import (
	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/db/repository"
)

type InviteRepository struct{}

func (r InviteRepository) Add(invite entity.Invite) error {
	return repository.Add(invite)
}

func (r InviteRepository) GetAll() []entity.Invite {
	return repository.GetAll[entity.Invite](entity.Invite{})
}

func (r InviteRepository) GetById(id int) (*entity.Invite, error) {
	return repository.GetById[entity.Invite](id, entity.Invite{})
}

func (r InviteRepository) Update(invite entity.Invite) error {
	return repository.Update(invite)
}

func (r InviteRepository) Delete(invite entity.Invite) error {
	return repository.Delete(invite)
}

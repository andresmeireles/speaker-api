package auth

import (
	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/db/repository"
)

type AuthRepository struct{}

func (a AuthRepository) Add(auth entity.Auth) error {
	return repository.Add(auth)
}

func (a AuthRepository) GetById(id int) (*entity.Auth, error) {
	return repository.GetById[entity.Auth](id, entity.Auth{})
}

func (a AuthRepository) GetAll() []entity.Auth {
	return repository.GetAll[entity.Auth](entity.Auth{})
}

func (a AuthRepository) Update(auth entity.Auth) error {
	return repository.Update(auth)
}

func (a AuthRepository) Delete(auth entity.Auth) error {
	return repository.Delete(auth)
}

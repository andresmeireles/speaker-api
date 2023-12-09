package config

import (
	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/db/repository"
)

type ConfigRepository struct{}

func (c ConfigRepository) GetAll() []entity.Config {
	return repository.GetAll[entity.Config](entity.Config{})
}

func (c ConfigRepository) GetById(id int) (*entity.Config, error) {
	return repository.GetById[entity.Config](id, entity.Config{})
}

func (c ConfigRepository) Add(config entity.Config) error {
	return repository.Add(config)
}

func (c ConfigRepository) Update(config entity.Config) error {
	return repository.Update(config)
}

func (c ConfigRepository) Delete(config entity.Config) error {
	return repository.Delete(config)
}

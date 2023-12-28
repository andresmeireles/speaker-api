package config

import (
	"github.com/andresmeireles/speaker/internal/db/entity"
)

func createConfig(name, value string, repository ConfigRepository) error {
	newConfig := entity.Config{
		Name:  name,
		Value: value,
	}

	if err := repository.Add(newConfig); err != nil {
		return err
	}

	return nil
}

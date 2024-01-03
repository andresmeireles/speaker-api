package config

import (
	"database/sql"

	"github.com/andresmeireles/speaker/internal/db/entity"
)

// Create or update a config
func Write(name, value string, repository ConfigRepository) error {
	config, err := repository.GetByName(name)

	if err == sql.ErrNoRows {
		return createConfig(name, value, repository)
	}

	if err != nil {
		return err
	}

	config.Value = value

	return repository.Update(*config)
}

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

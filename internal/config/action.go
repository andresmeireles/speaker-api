package config

import (
	"database/sql"

	"github.com/andresmeireles/speaker/internal/tools/servicelocator"
)

type Actions struct {
	configRepository ConfigRepository
}

func (a Actions) New(s servicelocator.ServiceLocator) any {
	return Actions{
		configRepository: servicelocator.Get[ConfigRepository](s),
	}
}

// Create or update a config.
func (a Actions) Write(name, value string) error {
	config, err := a.configRepository.GetByName(name)
	if err == sql.ErrNoRows {
		return createConfig(name, value, a.configRepository)
	}

	if err != nil {
		return err
	}

	config.Value = value

	return a.configRepository.Update(*config)
}

func createConfig(name, value string, repository ConfigRepository) error {
	newConfig := Config{
		Name:  name,
		Value: value,
	}

	if err := repository.Add(newConfig); err != nil {
		return err
	}

	return nil
}

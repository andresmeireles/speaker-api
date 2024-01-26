package config

import (
	"database/sql"
)

type Actions struct {
	configRepository ConfigRepository
}

func NewActions(configRepository ConfigRepository) Actions {
	return Actions{
		configRepository: configRepository,
	}
}

// Create or update a config.
func (a Actions) Write(name, value string) error {
	config, err := a.configRepository.GetByName(name)
	if err == sql.ErrNoRows {
		return a.createConfig(name, value)
	}

	if err != nil {
		return err
	}

	config.Value = value

	return a.configRepository.Update(*config)
}

func (a Actions) createConfig(name, value string) error {
	newConfig := Config{
		Name:  name,
		Value: value,
	}

	if err := a.configRepository.Add(newConfig); err != nil {
		return err
	}

	return nil
}

package config

import (
	"fmt"

	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/db/repository"
)

type ConfigRepository struct{}

func (c ConfigRepository) GetAll() ([]entity.Config, error) {
	rows, err := repository.GetAll[entity.Config]()

	if err != nil {
		return nil, err
	}

	configs := make([]entity.Config, 0)

	for rows.Next() {
		var config entity.Config
		if err := rows.Scan(&config.Id, &config.Name, &config.Value); err != nil {
			return nil, err
		}
		configs = append(configs, config)
	}

	return configs, nil
}

func (c ConfigRepository) GetById(id int) (*entity.Config, error) {
	row, err := repository.GetById[entity.Config](id)

	if err != nil {
		return nil, err
	}

	var config entity.Config

	if err := row.Scan(&config.Id, &config.Name, &config.Value); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c ConfigRepository) GetByName(name string) (*entity.Config, error) {
	config := new(entity.Config)
	query := fmt.Sprintf("SELECT * FROM %s WHERE name = $1 LIMIT 1", config.Table())
	result, err := repository.SingleQuery(query, name)

	if err != nil {
		return nil, result.Err()
	}

	if err := result.Scan(&config.Id, &config.Name, &config.Value); err != nil {
		return nil, err
	}

	return config, nil
}

func (c ConfigRepository) Add(config entity.Config) error {
	err := repository.Add(config)

	if err != nil && err.Error() == "pq: duplicate key value violates unique constraint \"configs_name_key\"" {
		return fmt.Errorf("config with name %s already exists", config.Name)
	}

	return err
}

func (c ConfigRepository) Update(config entity.Config) error {
	return repository.Update(config)
}

func (c ConfigRepository) Delete(config entity.Config) error {
	return repository.Delete(config)
}

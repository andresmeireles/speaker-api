package config

import (
	"fmt"

	"github.com/andresmeireles/speaker/internal/repository"
)

type ConfigRepository interface {
	GetAll() ([]Config, error)
	GetById(id int) (*Config, error)
	GetByName(name string) (*Config, error)
	Add(config Config) error
	Update(config Config) error
	Delete(config Config) error
}

type Repository struct {
	repository repository.RepositoryInterface
}

func NewRepository(repository repository.RepositoryInterface) ConfigRepository {
	return Repository{repository: repository}
}

func (c Repository) GetAll() ([]Config, error) {
	rows, err := c.repository.GetAll(Config{}.Table())

	if err != nil {
		return nil, err
	}

	configs := make([]Config, 0)

	for rows.Next() {
		var config Config
		if err := rows.Scan(&config.Id, &config.Name, &config.Value); err != nil {
			return nil, err
		}

		configs = append(configs, config)
	}

	return configs, nil
}

func (c Repository) GetById(id int) (*Config, error) {
	config := new(Config)
	row, err := c.repository.GetById(config.Table(), id)

	if err != nil {
		return nil, err
	}

	if err := row.Scan(&config.Id, &config.Name, &config.Value); err != nil {
		return nil, err
	}

	return config, nil
}

func (r Repository) GetByName(name string) (*Config, error) {
	config := new(Config)
	query := fmt.Sprintf("SELECT * FROM %s WHERE name = $1 LIMIT 1", config.Table())
	result, err := r.repository.SingleQuery(query, name)

	if err != nil {
		return nil, err
	}

	if err := result.Scan(&config.Id, &config.Name, &config.Value); err != nil {
		return nil, err
	}

	return config, nil
}

func (c Repository) Add(config Config) error {
	err := c.repository.Add(config)

	if err != nil && err.Error() == "pq: duplicate key value violates unique constraint \"configs_name_key\"" {
		return fmt.Errorf("config with name %s already exists", config.Name)
	}

	return err
}

func (c Repository) Update(config Config) error {
	return c.repository.Update(config)
}

func (c Repository) Delete(config Config) error {
	return c.repository.Delete(config)
}

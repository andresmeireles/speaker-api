package config

import (
	"fmt"

	"github.com/andresmeireles/speaker/internal/repository"
	"github.com/andresmeireles/speaker/internal/tools/servicelocator"
)

type ConfigRepository struct {
	repository repository.Repository[Config]
}

func (c ConfigRepository) New(s servicelocator.ServiceLocator) any {
	return ConfigRepository{
		repository: servicelocator.Get[repository.Repository[Config]](s),
	}
}

func (c ConfigRepository) GetAll() ([]Config, error) {
	rows, err := c.repository.GetAll()

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

func (c ConfigRepository) GetById(id int) (*Config, error) {
	row, err := c.repository.GetById(id)

	if err != nil {
		return nil, err
	}

	var config Config

	if err := row.Scan(&config.Id, &config.Name, &config.Value); err != nil {
		return nil, err
	}

	return &config, nil
}

func (r ConfigRepository) GetByName(name string) (*Config, error) {
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

func (c ConfigRepository) Add(config Config) error {
	err := c.repository.Add(config)

	if err != nil && err.Error() == "pq: duplicate key value violates unique constraint \"configs_name_key\"" {
		return fmt.Errorf("config with name %s already exists", config.Name)
	}

	return err
}

func (c ConfigRepository) Update(config Config) error {
	return c.repository.Update(config)
}

func (c ConfigRepository) Delete(config Config) error {
	return c.repository.Delete(config)
}

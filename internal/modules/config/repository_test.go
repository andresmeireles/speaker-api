package config_test

import (
	"testing"

	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/modules/config"
	"github.com/andresmeireles/speaker/testdata"
)

func TestMain(m *testing.M) {
	testdata.SetupDatabase(m)
}

func TestRepository(t *testing.T) {
	t.Run("should return configs", func(t *testing.T) {
		// arrange
		conf := entity.Config{
			Name:  "key",
			Value: "value",
		}
		conf2 := entity.Config{
			Name:  "key2",
			Value: "value2",
		}
		repository := config.ConfigRepository{}
		err := repository.Add(conf)

		if err != nil {
			t.Fatalf("expected nil, got %s", err)
		}

		err = repository.Add(conf2)

		if err != nil {
			t.Fatalf("expected nil, got %s", err)
		}

		// act
		configs, err := repository.GetAll()

		// assert
		if err != nil {
			t.Fatalf("expected nil, got %s", err)
		}

		if len(configs) != 2 {
			t.Fatalf("expected 2, got %d", len(configs))
		}
	})

	t.Run("should not save config with same name", func(t *testing.T) {
		// arrange
		conf := entity.Config{
			Name:  "key",
			Value: "value",
		}
		repository := config.ConfigRepository{}
		err := repository.Add(conf)

		if err != nil {
			t.Fatalf("expected nil, got %s", err)
		}

		conf2 := entity.Config{
			Name:  "key",
			Value: "value2",
		}

		// act
		err = repository.Add(conf2)

		// assert
		if err == nil {
			t.Fatalf("expected error, got nil")
		}

		if err.Error() != "config with name key already exists" {
			t.Fatalf("expected config with name key already exists, got %s", err.Error())
		}
	})

	t.Run("should return config by id", func(t *testing.T) {
		// arrange
		conf := entity.Config{
			Name:  "key",
			Value: "value",
		}
		repository := config.ConfigRepository{}
		err := repository.Add(conf)

		if err != nil {
			t.Fatalf("expected nil, got %s", err)
		}

		// act
		config, err := repository.GetById(1)

		// assert
		if err != nil {
			t.Fatalf("expected nil, got %s", err)
		}

		if config.Name != "key" {
			t.Fatalf("expected key, got %s", config.Name)
		}
	})
}

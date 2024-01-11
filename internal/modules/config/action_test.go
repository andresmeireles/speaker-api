package config_test

// import (
// 	"testing"

// 	"github.com/andresmeireles/speaker/internal/modules/config"
// )

// func TestActions(t *testing.T) {
// 	actions := config.NewActions()

// 	t.Run("should create a new config", func(t *testing.T) {
// 		// arrange
// 		repo := config.ConfigRepository{}

// 		// act
// 		err := actions.Write("key", "value")
// 		conf, getErr := repo.GetByName("key")

// 		// assert
// 		if err != nil {
// 			t.Errorf("unexpected error: %s", err)
// 		}

// 		if getErr != nil {
// 			t.Errorf("unexpected error: %s", getErr)
// 		}

// 		if conf.Name != "key" {
// 			t.Errorf("expected key, got %s", conf.Name)
// 		}

// 		if conf.Value != "value" {
// 			t.Errorf("expected value, got %s", conf.Value)
// 		}
// 	})

// 	t.Run("should update a config", func(t *testing.T) {
// 		// arrange
// 		repo := config.ConfigRepository{}
// 		_ = actions.Write("key", "value")

// 		// act
// 		err := actions.Write("key", "value2")
// 		conf, getErr := repo.GetByName("key")

// 		// assert
// 		if err != nil {
// 			t.Errorf("unexpected error: %s", err)
// 		}

// 		if getErr != nil {
// 			t.Errorf("unexpected error: %s", getErr)
// 		}

// 		if conf.Value != "value2" {
// 			t.Errorf("expected value2, got %s", conf.Value)
// 		}
// 	})
// }

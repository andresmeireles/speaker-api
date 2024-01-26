package env_test

import (
	"testing"

	"github.com/andresmeireles/speaker/internal/tools/env"
)

func TestEnv(t *testing.T) {
	t.Run("Get dev env as true", func(t *testing.T) {
		// assert
		t.Setenv("APP_MODE", "dev")

		// act
		result := env.IsDev()

		// assert
		if !result {
			t.Errorf("IsDev() = %v, want %v", result, true)
		}
	})

	t.Run("Get dev env as false", func(t *testing.T) {
		// assert
		t.Setenv("APP_MODE", "prod")

		// act
		result := env.IsDev()

		// assert
		if result {
			t.Errorf("IsDev() = %v, want %v", result, false)
		}
	})

	t.Run("Show dev as false when env not exists", func(t *testing.T) {
		// act
		result := env.IsDev()

		// assert
		if result {
			t.Errorf("IsDev() = %v, want %v", result, false)
		}
	})
}

package env_test

import (
	"errors"
	"os"
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

	t.Run("stack is true", func(t *testing.T) {
		// arrange
		os.Setenv("SHOW_STACK_TRACE", "true")
		defer os.Unsetenv("SHOW_STACK_TRACE")

		// act
		result := env.ShowStackTrace()

		// assert
		if !result {
			t.Fatal("Must be true")
		}
	})

	t.Run("stack is true", func(t *testing.T) {
		// act
		result := env.ShowStackTrace()

		// assert
		if result {
			t.Fatal("Must be false")
		}
	})

	t.Run("show error is true", func(t *testing.T) {
		// arrange
		os.Setenv("SHOW_ERROR_FILE", "true")

		// act
		result := env.ShowErrorFile()

		// assert
		if !result {
			t.Fatal("Must be true")
		}
	})

	t.Run("show error is false", func(t *testing.T) {
		// arrange
		os.Unsetenv("SHOW_ERROR_FILE")

		// act
		result := env.ShowErrorFile()

		// assert
		if result {
			t.Fatal("Must be false")
		}
	})

	t.Run("Should return AppKey", func(t *testing.T) {
		// arrange
		os.Setenv("APP_KEY", "app")
		defer os.Unsetenv("APP_KEY")

		// act
		app, err := env.AppKey()

		// asserts
		if err != nil {
			t.Fatalf("Err must be nil. Received %s", err)
		}

		if app != "app" {
			t.Fatalf("App key must be app. Received %s", app)
		}
	})

	t.Run("Should return error when app_key not exists", func(t *testing.T) {
		_, err := env.AppKey()

		// assert
		if err == nil {
			t.Fatal("Err must not be nil")
		}

		if !errors.Is(err, env.ErrAPPKeyNotExists) {
			t.Fatalf("Error is wrong. Received %s", err)
		}
	})
}

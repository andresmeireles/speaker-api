package logger_test

import (
	"log/slog"
	"os"
	"reflect"
	"testing"

	"github.com/andresmeireles/speaker/internal/tools/logger"
)

func TestLogger(t *testing.T) {
	t.Run("should return json logger", func(t *testing.T) {
		logger.Logger()

		// assert
		if reflect.TypeOf(slog.Default()).String() != "*slog.Logger" {
			t.Errorf("Logger() = given %v, want %v", reflect.TypeOf(slog.Default()).String(), "*slog.TextHandler")
		}
	})

	t.Run("should return json dev logger", func(t *testing.T) {
		os.Setenv("APP_MODE", "dev")

		logger.Logger()

		// assert
		if reflect.TypeOf(slog.Default()).String() != "*slog.Logger" {
			t.Errorf("Logger() = given %v, want %v", reflect.TypeOf(slog.Default()).String(), "*slog.TextHandler")
		}
	})
}

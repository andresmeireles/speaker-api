package logger_test

import (
	"log/slog"
	"reflect"
	"testing"

	"github.com/andresmeireles/speaker/internal/tools/logger"
)

func TestLogger(t *testing.T) {
	logger.Logger()

	// assert
	if reflect.TypeOf(slog.Default()).String() != "*slog.Logger" {
		t.Errorf("Logger() = given %v, want %v", reflect.TypeOf(slog.Default()).String(), "*slog.TextHandler")
	}
}

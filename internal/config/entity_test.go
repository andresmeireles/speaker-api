package config_test

import (
	"testing"

	"github.com/andresmeireles/speaker/internal/config"
)

func TestConfigTable(t *testing.T) {
	config := config.Config{}

	if config.Table() != "configs" {
		t.Fatalf("expected configs, got %s", config.Table())
	}

	if config.GetId() != 0 {
		t.Fatalf("expected 0, got %d", config.GetId())
	}
}

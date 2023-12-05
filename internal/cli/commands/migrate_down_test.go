package commands_test

import (
	"testing"

	"github.com/andresmeireles/speaker/internal/cli/commands"
)

func TestMigrateDown(t *testing.T) {
	// act
	err := commands.MigrateDown().Execute()

	// assert
	if err != nil {
		t.Fatalf("expected nil, got %s", err)
	}
}

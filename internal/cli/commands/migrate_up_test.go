package commands_test

import (
	"testing"

	"github.com/andresmeireles/speaker/internal/cli/commands"
	"github.com/andresmeireles/speaker/testdata"
)

func TestMain(m *testing.M) {
	testdata.SetupDatabase(m)
}

func TestMigrateUp(t *testing.T) {
	// act
	err := commands.MigrateUp().Execute()

	// assert
	if err != nil {
		t.Fatalf("expected nil, got %s", err)
	}
}

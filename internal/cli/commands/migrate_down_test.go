package commands_test

import (
	"testing"

	"github.com/andresmeireles/speaker/internal/cli/commands"
	"github.com/andresmeireles/speaker/testdata/mocks"
)

func TestMigrateDown(t *testing.T) {
	// arrange
	migration := mocks.Migrationauxcmd{}
	migration.EXPECT().Down().Return(nil)

	// act
	err := commands.MigrateDown(&migration).Execute()

	// assert
	if err != nil {
		t.Fatalf("expected nil, got %s", err)
	}
}

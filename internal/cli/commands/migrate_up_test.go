package commands_test

import (
	"testing"

	"github.com/andresmeireles/speaker/internal/cli/commands"
	"github.com/andresmeireles/speaker/testdata/mocks"
)

func TestMigrateUp(t *testing.T) {
	// arrange
	migration := mocks.Migration{}
	migration.EXPECT().Up().Return(nil)

	// act
	err := commands.MigrateUp(&migration).Execute()

	// assert
	if err != nil {
		t.Fatalf("expected nil, got %s", err)
	}
}

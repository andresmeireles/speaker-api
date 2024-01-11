package commands_test

import (
	"testing"

	"github.com/andresmeireles/speaker/internal/cli/commands"
)

func TestListUser(t *testing.T) {
	// arrange
	command := commands.ListUser()

	// act
	err := command.Execute()

	// assert
	if err != nil {
		t.Fatalf("expected nil, got %s", err)
	}
}

package commands_test

import (
	"testing"

	"github.com/andresmeireles/speaker/internal/cli/commands"
)

func TestSetAppKey(t *testing.T) {
	// arrange
	command := commands.SetAppKey()
	command.SetArgs([]string{"123"})

	// act
	err := command.Execute()

	// assert
	if err != nil {
		t.Fatalf("expected nil, got %s", err)
	}
}

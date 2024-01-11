package commands_test

import (
	"testing"

	"github.com/andresmeireles/speaker/internal/cli/commands"
)

func TestCreateUser(t *testing.T) {
	t.Run("create user", func(t *testing.T) {
		// arrange
		command := commands.CreateUser()
		command.SetArgs([]string{"-n", "Person 1", "-e", "123"})

		// act
		err := command.Execute()

		// assert
		if err != nil {
			t.Fatalf("expected nil, got %s", err)
		}
	})
}

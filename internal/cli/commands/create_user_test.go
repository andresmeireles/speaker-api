package commands_test

import (
	"testing"

	"github.com/andresmeireles/speaker/internal/cli/commands"
	"github.com/andresmeireles/speaker/internal/user"
	"github.com/andresmeireles/speaker/testdata/mocks"
)

func TestCreateUser(t *testing.T) {
	t.Run("create user", func(t *testing.T) {
		// arrange
		userRepository := mocks.UserRepository{}
		userRepository.EXPECT().Add(user.User{Name: "Person 1", Email: "123"}).Return(nil)
		command := commands.CreateUser(&userRepository)
		command.SetArgs([]string{"-n", "Person 1", "-e", "123"})

		// act
		err := command.Execute()

		// assert
		if err != nil {
			t.Fatalf("expected nil, got %s", err)
		}
	})
}

package codesender_test

import (
	"testing"

	"github.com/andresmeireles/speaker/internal/codesender"
	"github.com/andresmeireles/speaker/internal/user"
	"github.com/andresmeireles/speaker/testdata"
)

func TestMain(m *testing.M) {
	testdata.SetupDatabase(m)
}

func TestSendCode(t *testing.T) {
	codesender := testdata.GetService[codesender.Actions]()
	userRepo := testdata.GetService[user.UserRepository]()

	t.Run("should send code", func(t *testing.T) {
		// arrange
		user := user.User{
			Email: "someemail",
		}
		userRepo.Add(user)
		user.Id = 1

		// act
		code, err := codesender.CreateCode(user)

		// assert
		if err != nil {
			t.Error(err)
		}

		if code == "" {
			t.Error("code should not be empty")
		}
	})
}

package codesender_test

import (
	"testing"
	"time"

	"github.com/andresmeireles/speaker/internal/codesender"
	"github.com/andresmeireles/speaker/internal/user"
	"github.com/andresmeireles/speaker/testdata"
	"github.com/stretchr/testify/mock"
)

func TestSendCode(t *testing.T) {
	authCodeRepo := new(testdata.AuthCodeRepositoryMock)
	action := codesender.NewAction(authCodeRepo)

	t.Run("should send code", func(t *testing.T) {
		// arrange
		user := user.User{}
		authCodeRepo.On(
			"Add",
			mock.MatchedBy(func(authCode codesender.AuthCode) bool { return true }),
		).Return(nil)

		// act
		code, err := action.CreateCode(user)

		// assert
		if err != nil {
			t.Error(err)
		}

		if code == "" {
			t.Error("code should not be empty")
		}
	})

	t.Run("verify code", func(t *testing.T) {
		// arrange
		email := "user email"
		authCode := codesender.AuthCode{
			User:      user.User{Email: email},
			ExpiresAt: time.Now().Add(time.Minute * 5),
		}
		authCodeRepo.On(
			"GetByCode",
			mock.MatchedBy(func(code string) bool { return true }),
		).Return(authCode, nil)

		// act
		err := action.VerifyCode(email, "code")

		// assert
		if err != nil {
			t.Error(err)
		}
	})
}

package auth_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/andresmeireles/speaker/internal/auth"
	"github.com/andresmeireles/speaker/internal/user"
	"github.com/andresmeireles/speaker/testdata/mocks"
)

var errTest = fmt.Errorf("test error")

func TestMain(m *testing.M) {
	os.Setenv("APP_KEY", "key")
	m.Run()
	os.Unsetenv("APP_KEY")
}

//nolint:funlen
func TestActions(t *testing.T) {
	repositoryMock := mocks.Repositoryauth{}
	emailMock := mocks.Etools{}
	codeSenderActionMock := mocks.Servicecodesender{}
	userRepositoryMock := mocks.UserRepositoryuser{}

	authAction := auth.NewAction(
		&repositoryMock,
		&userRepositoryMock,
		&emailMock,
		&codeSenderActionMock,
	)

	t.Run("should create jwt", func(t *testing.T) {
		// arrange
		user := user.User{
			Id:    1,
			Name:  "John",
			Email: "john.doe@email.com",
		}
		token, _ := authAction.CreateToken("andre.meireles", user.Email, time.Hour*24)
		repositoryMock.EXPECT().
			Add(auth.Auth{User: user, UserId: 1, Hash: token}).
			Return(nil)

		// act
		result, err := authAction.CreateJWT(user, false)

		// assert
		if err != nil {
			t.Fatalf("Must be nil. Received: %s", err)
		}

		if result.UserId != 1 {
			t.Fatalf("User id must be 1. Received: %d", result.UserId)
		}
	})

	t.Run("should create jwt for two weeks", func(t *testing.T) {
		// arrange
		user := user.User{
			Id:    1,
			Name:  "John",
			Email: "john.doe@email.com",
		}
		token, _ := authAction.CreateToken("andre.meireles", user.Email, time.Hour*24*7*2)
		repositoryMock.EXPECT().
			Add(auth.Auth{User: user, UserId: 1, Hash: token}).
			Return(nil)

		// act
		result, err := authAction.CreateJWT(user, true)

		// assert
		if err != nil {
			t.Fatalf("Must be nil. Received: %s", err)
		}

		if result.UserId != 1 {
			t.Fatalf("User id must be 1. Received: %d", result.UserId)
		}
	})

	t.Run("should validate jwt", func(t *testing.T) {
		// arrange
		token, _ := authAction.CreateToken("andre.meireles", "john.doe@email.com", time.Hour*24)

		// act
		result := authAction.ValidateJwt(token)

		// assert
		if !result {
			t.Fatal("Must be true")
		}
	})

	t.Run("Should return false with invalid token", func(t *testing.T) {
		// arrange
		token := "invalid"

		// act
		result := authAction.ValidateJwt(token)

		// assert
		if result {
			t.Fatal("Must be false")
		}
	})

	t.Run("Should return false for expired token", func(t *testing.T) {
		// arrange
		token, _ := authAction.CreateToken("andre.meireles", "john.doe@email.com", time.Hour*-24)

		// act
		result := authAction.ValidateJwt(token)

		// assert
		if result {
			t.Fatal("Must be false")
		}
	})

	t.Run("Should return false when email not exists", func(t *testing.T) {
		// arrange
		userRepositoryMock.EXPECT().
			GetByEmail("john.doe@email.com").
			Return(user.User{}, errTest)

		// act
		result := authAction.HasEmail("john.doe@email.com")

		// assert
		if result {
			t.Fatal("Must be false")
		}
	})

	t.Run("Should return true when email exists", func(t *testing.T) {
		// arrange
		userRepositoryMock.EXPECT().
			GetByEmail("john.doe@email.com").
			Return(user.User{}, nil)

		// act
		result := authAction.HasEmail("john.doe@email.com")

		// assert
		if !result {
			t.Fatal("Must be true")
		}
	})

	t.Run("Should return no error when send a code", func(t *testing.T) {
		// arrange
		user := user.User{
			Id:    1,
			Name:  "John",
			Email: "john.doe@email.com",
		}
		codeSenderActionMock.EXPECT().
			CreateCode(user).
			Return("code", nil)

		userRepositoryMock.EXPECT().
			GetByEmail(user.Email).
			Return(user, nil)

		emailMock.EXPECT().Send("code", "john.doe@email.com").Return(nil)

		// act
		err := authAction.SendCode("john.doe@email.com")

		// assert
		if err != nil {
			t.Fatalf("Must be nil. Received: %s", err)
		}
	})
}

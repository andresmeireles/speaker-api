package auth_test

import (
	"errors"
	"testing"

	"github.com/andresmeireles/speaker/internal/auth"
	"github.com/andresmeireles/speaker/testdata/mocks"
)

var ErrRepo = errors.New("repo err")

func createTestSubjects() (auth.AuthRepository, *mocks.RepositoryInterfacerepository) {
	repository := mocks.RepositoryInterfacerepository{}
	authRepository := auth.NewRepository(&repository)

	return authRepository, &repository
}

func TestRepository(t *testing.T) {
	authRepository, repository := createTestSubjects()

	t.Run("should update", func(t *testing.T) {
		// arrange
		auth := auth.Auth{}
		repository.EXPECT().Update(auth).Return(nil)

		// act
		err := authRepository.Update(auth)

		// assert
		if err != nil {
			t.Fatal("Must return nil")
		}
	})

	t.Run("should delete", func(t *testing.T) {
		// arrange
		auth := auth.Auth{}
		repository.EXPECT().Delete(auth).Return(nil)

		// act
		err := authRepository.Delete(auth)

		// assert
		if err != nil {
			t.Fatal("Err must be nil")
		}
	})

	t.Run("should delete with error", func(t *testing.T) {
		// arrange
		a, r := createTestSubjects()
		auth := auth.Auth{}

		r.EXPECT().Delete(auth).Return(ErrRepo)

		// act
		err := a.Delete(auth)

		// assert
		if err == nil {
			t.Fatal("Err must not be nil")
		}

		if !errors.Is(ErrRepo, err) {
			t.Fatalf("error must be %s", ErrRepo)
		}
	})
}

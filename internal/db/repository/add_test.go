package repository_test

import (
	"testing"

	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/db/repository"
	"github.com/andresmeireles/speaker/testdata"
)

func TestMain(m *testing.M) {
	testdata.SetupDatabase(m)
}

func TestAdd(t *testing.T) {
	t.Run("add register", func(t *testing.T) {
		somePerson := entity.Person{
			Name: "Some Person",
		}

		err := repository.Add(somePerson)

		if err != nil {
			t.Fatalf("expected nil, got %s", err)
		}
	})
}

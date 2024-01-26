package repository_test

import (
	"os"
	"testing"

	"github.com/andresmeireles/speaker/internal/db"
	"github.com/andresmeireles/speaker/internal/person"
	"github.com/andresmeireles/speaker/internal/repository"
)

func TestMain(m *testing.M) {
	os.Setenv("DB_DRIVER", "sqlite3")
}

func TestAdd(t *testing.T) {
	r := repository.NewRepository(
		db.Connection{},
	)

	t.Run("add register", func(t *testing.T) {
		// arrange
		somePerson := person.Person{
			Name: "Some Person",
		}

		// act
		err := r.Add(somePerson)

		if err != nil {
			t.Fatalf("expected nil, got %s", err)
		}
	})
}

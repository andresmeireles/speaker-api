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
	r := testdata.GetService[repository.Repository[entity.Person]]()

	t.Run("add register", func(t *testing.T) {
		// arrange
		somePerson := entity.Person{
			Name: "Some Person",
		}

		// act
		err := r.Add(somePerson)

		if err != nil {
			t.Fatalf("expected nil, got %s", err)
		}
	})
}

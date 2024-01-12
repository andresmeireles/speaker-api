package repository_test

import (
	"testing"

	"github.com/andresmeireles/speaker/internal/person"
	"github.com/andresmeireles/speaker/internal/repository"
	"github.com/andresmeireles/speaker/testdata"
)

func TestMain(m *testing.M) {
	testdata.SetupDatabase(m)
}

func TestAdd(t *testing.T) {
	r := testdata.GetService[repository.Repository]()

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

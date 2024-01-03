package person_test

import (
	"testing"

	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/modules/person"
	"github.com/andresmeireles/speaker/testdata"
)

func TestMain(m *testing.M) {
	testdata.SetupDatabase(m)
}

func TestUpdate(t *testing.T) {
	t.Run("should update person name", func(t *testing.T) {
		// arrange
		repo := person.PersonRepository{}
		p1 := entity.Person{
			Name: "Andre",
		}
		repo.Add(p1)
		p2 := entity.Person{
			Name: "Yasmim",
		}
		repo.Add(p2)

		// act
		dbP1, err := repo.GetById(1)

		if err != nil {
			t.Fatalf("expected nil, got %s", err)
		}

		dbP1.Name = "André Meireles"
		err = repo.Update(*dbP1)

		// assert
		if err != nil {
			t.Fatalf("Expected nil, got %s", err)
		}

		if dbP1.Name != "André Meireles" {
			t.Fatalf("Expected André Meireles, got %s", dbP1.Name)
		}
	})
}

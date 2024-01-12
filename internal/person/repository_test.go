package person_test

import (
	"testing"

	"github.com/andresmeireles/speaker/internal/person"
	"github.com/andresmeireles/speaker/internal/repository"
	"github.com/andresmeireles/speaker/testdata"
)

func TestMain(m *testing.M) {
	testdata.SetupDatabase(m)
}

func cleanDb() {
	r := testdata.GetService[repository.Repository[person.Person]]()
	r.Query("DELETE FROM persons")
}

func TestUpdate(t *testing.T) {
	t.Run("should update person name", func(t *testing.T) {
		// arrange
		repo := testdata.GetService[person.PersonRepository]()
		p1 := person.Person{
			Name: "Andre",
		}
		repo.Add(p1)
		p2 := person.Person{
			Name: "Yasmim",
		}
		repo.Add(p2)
		p, err := repo.GetByName("Andre")

		if err != nil {
			t.Fatalf("expected nil, got %s", err)
		}

		// act
		dbP1, err := repo.GetById(p.Id)

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

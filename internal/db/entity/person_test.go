package entity_test

import (
	"testing"

	"github.com/andresmeireles/speaker/internal/db/entity"
)

func TestPersonEntity(t *testing.T) {
	t.Run("Get table name", func(t *testing.T) {
		person := entity.Person{}

		if person.Table() != "persons" {
			t.Fatalf("expected persons, got %s", person.Table())
		}
	})

	t.Run("Get id", func(t *testing.T) {
		person := entity.Person{}
		if person.GetId() != 0 {
			t.Fatalf("expected 0, got %d", person.GetId())
		}
	})
}

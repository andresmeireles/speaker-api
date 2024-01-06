package web_test

import (
	"strings"
	"testing"

	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/web"
)

func TestDecodePostBody(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// arrange
		body := strings.NewReader(`{"id": 1, "name": "Person 1", "FIELD": "VALUE"}`)

		// act
		person, err := web.DecodePostBody[entity.Person](body)

		// assert
		if err != nil {
			t.Fatalf("expected nil, got %s", err)
		}
		if person.Name != "Person 1" {
			t.Fatalf("expected Person 1, got %s", person.Name)
		}
		if person.GetId() != 1 {
			t.Fatalf("expected 1, got %d", person.GetId())
		}
	})

	t.Run("Test without complete body", func(t *testing.T) {
		// arrange
		bodyWithId := strings.NewReader(`{"id": 1}`)
		bodyWithName := strings.NewReader(`{"name": "Person 1"}`)

		// act
		personWithId, errWithId := web.DecodePostBody[entity.Person](bodyWithId)
		personWithName, errWithName := web.DecodePostBody[entity.Person](bodyWithName)

		// assert
		if errWithId != nil || errWithName != nil {
			t.Fatalf("expected nil, got %s as errWithId and %s as errWithName", errWithId, errWithName)
		}
		if personWithId.GetId() != 1 && personWithId.Name == "" {
			t.Fatalf("expected id 1 and no name, got %d and name %s", personWithId.GetId(), personWithId.Name)
		}
		if personWithName.Name != "Person 1" && personWithName.GetId() != 0 {
			t.Fatalf("expected Person 1 and id 0, got %s and id %d", personWithName.Name, personWithName.GetId())

		}
	})
}

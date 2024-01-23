package invite_test

import (
	"testing"
	"time"

	"github.com/andresmeireles/speaker/internal/invite"
	"github.com/andresmeireles/speaker/internal/person"
)

func TestInvite(t *testing.T) {
	t.Run("should return correct table name", func(t *testing.T) {
		invite := invite.Invite{}

		if invite.Table() != "invites" {
			t.Fatalf("expected invites, got %s", invite.Table())
		}
	})
	t.Run("should correct return map", func(t *testing.T) {
		// arrange
		invite := invite.Invite{
			PersonId: 1,
			Person: person.Person{
				Name: "Person 1",
			},
			Theme: "Theme",
			Time:  1,
			Date:  time.Now(),
		}

		// act
		json := invite.ToJson()

		// assert
		if json["person_id"] != 1 {
			t.Fatalf("expected Person 1, got %s", json["person"])
		}

		if json["theme"] != "Theme" {
			t.Fatalf("expected Theme, got %s", json["theme"])
		}

		if json["time"] != 1 {
			t.Fatalf("expected 1, got %d", json["time"])
		}
	})
	t.Run("should return correct id", func(t *testing.T) {
		// act
		invite := invite.Invite{}

		// assert
		if invite.GetId() != 0 {
			t.Fatalf("expected 0, got %d", invite.GetId())
		}
	})
}

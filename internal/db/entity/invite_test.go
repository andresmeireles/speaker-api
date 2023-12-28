package entity_test

// import (
// 	"testing"

// 	"github.com/andresmeireles/speaker/internal/db/entity"
// )

// func TestInviteTable(t *testing.T) {
// 	invite := entity.Invite{}

// 	if invite.Table() != "invites" {
// 		t.Fatalf("expected invites, got %s", invite.Table())
// 	}
// }

// func TestInviteToJson(t *testing.T) {
// 	// arrange
// 	invite := entity.Invite{
// 		Person: entity.Person{
// 			Name: "Person 1",
// 		},
// 		Theme:      "Theme",
// 		Time:       1,
// 		Date:       2,
// 		Accepted:   true,
// 		Remembered: true,
// 	}

// 	// act
// 	json := invite.ToJson()

// 	// assert
// 	if json["person"] != 0 {
// 		t.Fatalf("expected Person 1, got %s", json["person"])
// 	}
// 	if json["theme"] != "Theme" {
// 		t.Fatalf("expected Theme, got %s", json["theme"])
// 	}
// 	if json["time"] != 1 {
// 		t.Fatalf("expected 1, got %d", json["time"])
// 	}
// }

// func TestInviteGetId(t *testing.T) {
// 	// act
// 	invite := entity.Invite{}

// 	// assert
// 	if invite.GetId() != 0 {
// 		t.Fatalf("expected 0, got %d", invite.GetId())
// 	}
// }

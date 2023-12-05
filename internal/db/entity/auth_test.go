package entity_test

import (
	"testing"

	"github.com/andresmeireles/speaker/internal/db/entity"
)

func TestGetId(t *testing.T) {
	auth := entity.Auth{
		User: entity.User{},
		Hash: "hash",
	}

	if auth.GetId() != 0 {
		t.Fatalf("expected 0, got %d", auth.GetId())
	}
}

func TestToJson(t *testing.T) {
	// arrange
	auth := entity.Auth{
		User: entity.User{
			Name:      "Person 1",
			Telephone: "123",
		},
		Hash: "hash",
	}

	// act
	json := auth.ToJson()

	// assert
	if json["hash"] != "hash" {
		t.Fatalf("expected hash, got %s", json["hash"])
	}
}

func TestTable(t *testing.T) {
	auth := entity.Auth{
		User: entity.User{},
		Hash: "hash",
	}

	if auth.Table() != "auths" {
		t.Fatalf("expected auths, got %s", auth.Table())
	}
}

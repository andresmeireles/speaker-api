package auth_test

import (
	"testing"

	"github.com/andresmeireles/speaker/internal/auth"
	"github.com/andresmeireles/speaker/internal/user"
)

func TestGetId(t *testing.T) {
	auth := auth.Auth{
		User: user.User{},
		Hash: "hash",
	}

	if auth.GetId() != 0 {
		t.Fatalf("expected 0, got %d", auth.GetId())
	}
}

func TestToJson(t *testing.T) {
	// arrange
	auth := auth.Auth{
		User: user.User{
			Name:  "Person 1",
			Email: "123",
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
	auth := auth.Auth{
		User: user.User{},
		Hash: "hash",
	}

	if auth.Table() != "auths" {
		t.Fatalf("expected auths, got %s", auth.Table())
	}
}

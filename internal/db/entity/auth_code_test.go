package entity_test

import (
	"testing"
	"time"

	"github.com/andresmeireles/speaker/internal/db/entity"
)

func TestAuthCode(t *testing.T) {
	auth := entity.AuthCode{
		UserId: 1,
		User: entity.User{
			Name:  "Person 1",
			Email: "123",
		},
		Code:      "code",
		ExpiresAt: time.Now(),
	}

	t.Run("Get table name", func(t *testing.T) {
		if auth.Table() != "auth_codes" {
			t.Fatalf("expected auth_codes, got %s", auth.Table())
		}
	})

	t.Run("Get id", func(t *testing.T) {
		if auth.GetId() != 0 {
			t.Fatalf("expected 0, got %d", auth.GetId())
		}
	})

	t.Run("Expired is false", func(t *testing.T) {
		auth.ExpiresAt = time.Now().Add(time.Minute * 5)
		if auth.IsExpired() {
			t.Fatalf("expected false, got true")
		}
	})

	t.Run("Expired is true", func(t *testing.T) {
		auth.ExpiresAt = time.Now().Add(time.Minute * -5)
		if !auth.IsExpired() {
			t.Fatalf("expected true, got false")
		}
	})
}

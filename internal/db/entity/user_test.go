package entity_test

import (
	"testing"

	"github.com/andresmeireles/speaker/internal/db/entity"
)

func TestUserEntity(t *testing.T) {
	t.Run("Get table name", func(t *testing.T) {
		user := entity.User{}
		if user.Table() != "users" {
			t.Fatalf("expected users, got %s", user.Table())
		}
	})
	t.Run("Get id", func(t *testing.T) {
		user := entity.User{}
		if user.GetId() != 0 {
			t.Fatalf("expected 0, got %d", user.GetId())
		}
	})
}

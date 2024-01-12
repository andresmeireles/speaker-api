package user_test

import (
	"testing"

	"github.com/andresmeireles/speaker/internal/user"
)

func TestUserEntity(t *testing.T) {
	t.Run("Get table name", func(t *testing.T) {
		u := user.User{}
		if u.Table() != "users" {
			t.Fatalf("expected users, got %s", u.Table())
		}
	})
	t.Run("Get id", func(t *testing.T) {
		u := user.User{}
		if u.GetId() != 0 {
			t.Fatalf("expected 0, got %d", u.GetId())
		}
	})
}

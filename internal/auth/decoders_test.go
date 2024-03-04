package auth_test

import (
	"testing"

	"github.com/andresmeireles/speaker/internal/auth"
)

func TestDecoders(t *testing.T) {
	t.Run("create an email form struct", func(t *testing.T) {
		s := auth.EmailForm{"john.doe@email.com"}

		if s.Email != "john.doe@email.com" {
			t.Fatal("Email must be john.doe@email.com")
		}
	})

	t.Run("create a code form struct", func(t *testing.T) {
		s := auth.CodeForm{"c", "e@m.com", false}

		if s.Code != "c" {
			t.Fatalf("Code must be c. Received %s", s.Code)
		}

		if s.Email != "e@m.com" {
			t.Fatalf("Email must be e@m.com. Received %s", s.Email)
		}

		if s.Remember {
			t.Fatal("Remember must be false")
		}
	})
}

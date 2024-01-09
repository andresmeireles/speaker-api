package repository_test

import (
	"testing"
	"time"

	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/db/repository"
)

func TestDelete(t *testing.T) {
	t.Run("delete register", func(t *testing.T) {
		// arrange
		user := entity.User{
			Name:  "Person 1",
			Email: "123",
		}
		repository.Add(user)

		// act
		err := repository.Delete(user)
		allUsers, gErr := repository.GetAll[entity.User]()

		// assert
		if err != nil {
			t.Fatalf("expected nil, got %s", err)
		}

		if gErr != nil {
			t.Fatalf("expected nil, got %s", gErr)
		}

		numberOfRegisters := 0
		for allUsers.Next() {
			numberOfRegisters++
		}

		if numberOfRegisters != 0 {
			t.Fatalf("expected 0, got %d", numberOfRegisters)
		}
	})

	t.Run("Test Delete With Relationships", func(t *testing.T) {
		// arrange
		person := entity.Person{
			Name: "Person 1",
		}
		invite := entity.Invite{
			PersonId:   1,
			Theme:      "Theme",
			Time:       1,
			Date:       time.Now(),
			Accepted:   true,
			Remembered: true,
		}
		repository.Add(person)
		repository.Add(invite)

		// act
		err := repository.Delete(person)

		// assert
		if err != nil {
			t.Fatalf("expected nil, got %s", err)
		}
	})
}

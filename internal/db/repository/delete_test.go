package repository_test

import (
	"testing"
	"time"

	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/db/repository"
)

func TestDelete(t *testing.T) {
	// arrange
	user := entity.User{
		Name:  "Person 1",
		Email: "123",
	}
	repository.Add(user)

	// act
	err := repository.Delete(user)
	allUsers := repository.GetAll[entity.User](entity.User{})

	// assert
	if err != nil {
		t.Fatalf("expected nil, got %s", err)
	}
	if len(allUsers) != 0 {
		t.Fatalf("expected 0, got %d", len(allUsers))
	}
}

func TestDeleteWithRelationships(t *testing.T) {
	// arrange
	person := entity.Person{
		Name: "Person 1",
	}
	invite := entity.Invite{
		Person:     person,
		Theme:      "Theme",
		Time:       1,
		Date:       int(time.Now().Unix()),
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
}

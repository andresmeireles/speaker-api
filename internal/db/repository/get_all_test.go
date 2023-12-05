package repository_test

import (
	"testing"

	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/db/repository"
)

func TestGetAll(t *testing.T) {
	// arrange
	person1 := entity.Person{
		Name: "Person 1",
	}
	person2 := entity.Person{
		Name: "Person 2",
	}

	err := repository.Add(person1)

	if err != nil {
		t.Fatalf("expected nil, got %s", err)
	}

	err = repository.Add(person2)

	if err != nil {
		t.Fatalf("expected nil, got %s", err)
	}

	// act
	people := repository.GetAll[entity.Person](person1)

	// assert
	if len(people) != 2 {
		t.Fatalf("expected 2, got %d", len(people))
	}
}

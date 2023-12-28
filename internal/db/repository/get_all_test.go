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
	people, err := repository.GetAll[entity.Person]()

	// assert
	if err != nil {
		t.Fatalf("expected nil, got %s", err)
	}

	numberOfRegisters := 0

	for people.Next() {
		numberOfRegisters++
	}

	if numberOfRegisters != 2 {
		t.Fatalf("expected 2, got %d", numberOfRegisters)
	}
}

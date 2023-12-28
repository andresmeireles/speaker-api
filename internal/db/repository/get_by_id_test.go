package repository_test

// import (
// 	"testing"

// 	"github.com/andresmeireles/speaker/internal/db/entity"
// 	"github.com/andresmeireles/speaker/internal/db/repository"
// )

// func TestGetById(t *testing.T) {
// 	// arrange
// 	person := entity.Person{
// 		Name: "Person 1",
// 	}

// 	repository.Add(person)

// 	// act
// 	row := repository.GetById[entity.Person](1)

// 	entity := new(entity.Person)

// 	if err := row.Scan(entity.Id, entity.Name); err != nil {
// 		t.Fatalf("expected nil, got %s", err)
// 	}

// 	// assert
// 	if entity.GetId() != 1 {
// 		t.Fatalf("expected 1, got %d", entity.GetId())
// 	}
// 	if entity.Name != "Person 1" {
// 		t.Fatalf("expected Person 1, got %s", entity.Name)
// 	}
// }

// func TestBiggerEntity(t *testing.T) {
// 	// arrange
// 	u := entity.User{
// 		Name:  "Person 1",
// 		Email: "123",
// 	}
// 	repository.Add(u)

// 	// act
// 	row := repository.GetById[entity.User](1)

// 	// assert
// 	if user.GetId() != 1 {
// 		t.Fatalf("expected 1, got %d", user.GetId())
// 	}
// 	if user.Name != "Person 1" {
// 		t.Fatalf("expected Person 1, got %s", user.Name)
// 	}
// 	if user.Email != "123" {
// 		t.Fatalf("expected 123, got %s", user.Email)
// 	}
// }

// func TestNotFound(t *testing.T) {
// 	// arrange
// 	person := entity.Person{
// 		Name: "Person 1",
// 	}
// 	repository.Add(person)

// 	// act
// 	res, _ := repository.GetById[entity.Person](2, person)

// 	// assert
// 	if res != nil {
// 		t.Fatalf("expected error, got nil")
// 	}
// }

// func TestCallForthElement(t *testing.T) {
// 	// arrange
// 	person := entity.Person{
// 		Name: "Person",
// 	}
// 	repository.Add(person)
// 	person.Name = "Person 1"
// 	repository.Add(person)
// 	person.Name = "Person 2"
// 	repository.Add(person)
// 	person.Name = "Person 3"
// 	repository.Add(person)

// 	// act
// 	res, _ := repository.GetById[entity.Person](3, person)

// 	// assert
// 	if res.Name != "Person 3" {
// 		t.Fatalf("expected Person 3, got %s", res.Name)
// 	}
// 	if res.GetId() != 3 {
// 		t.Fatalf("expected 3, got %d", res.GetId())
// 	}
// }

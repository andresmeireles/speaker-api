package repository_test

// import (
// 	"testing"

// 	"github.com/andresmeireles/speaker/internal/db/entity"
// 	"github.com/andresmeireles/speaker/internal/repository"
// )

// func TestUpdate(t *testing.T) {
// 	// arrange
// 	person := entity.Person{
// 		Name: "Person 1",
// 	}
// 	repository.Add(person)

// 	// act
// 	person.Name = "Person 2"
// 	err := repository.Update[entity.Person](person)
// 	personFromDB, _ := repository.GetById[entity.Person](1, person)

// 	// arrange
// 	if err != nil {
// 		t.Fatalf("expected nil, got %s", err)
// 	}
// 	if personFromDB.Name != "Person 2" {
// 		t.Fatalf("expected Person 2, got %s", personFromDB.Name)
// 	}
// }

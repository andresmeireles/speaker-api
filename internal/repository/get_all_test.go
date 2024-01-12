package repository_test

// import (
// 	"testing"

// 	"github.com/andresmeireles/speaker/internal/db"
// 	"github.com/andresmeireles/speaker/internal/db/entity"
// 	"github.com/andresmeireles/speaker/internal/repository"
// )

// func TestGetAll(t *testing.T) {
// 	// arrange
// 	d, _ := db.GetDB()
// 	r := repository.Repository[entity.Person]{
// 		Db: d,
// 	}
// 	person1 := entity.Person{
// 		Name: "Person 1",
// 	}
// 	person2 := entity.Person{
// 		Name: "Person 2",
// 	}

// 	err := r.Add(person1)

// 	if err != nil {
// 		t.Fatalf("expected nil, got %s", err)
// 	}

// 	err = r.Add(person2)

// 	if err != nil {
// 		t.Fatalf("expected nil, got %s", err)
// 	}

// 	// act
// 	people, err := r.GetAll()

// 	// assert
// 	if err != nil {
// 		t.Fatalf("expected nil, got %s", err)
// 	}

// 	numberOfRegisters := 0

// 	for people.Next() {
// 		numberOfRegisters++
// 	}

// 	if numberOfRegisters != 2 {
// 		t.Fatalf("expected 2, got %d", numberOfRegisters)
// 	}
// }

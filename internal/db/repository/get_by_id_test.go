package repository_test

// import (
// 	"testing"

// 	"github.com/andresmeireles/speaker/internal/db"
// 	"github.com/andresmeireles/speaker/internal/db/entity"
// 	"github.com/andresmeireles/speaker/internal/db/repository"
// )

// const (
// 	PERSON_NAME = "Person 1"
// )

// func TestGetById(t *testing.T) {
// 	d, _ := db.GetDB()
// 	r := repository.Repository[entity.Person]{
// 		Db: d,
// 	}
// 	ru := repository.Repository[entity.User]{
// 		Db: d,
// 	}
// 	t.Run("get by id", func(t *testing.T) {
// 		// arrange
// 		person := entity.Person{
// 			Name: PERSON_NAME,
// 		}

// 		r.Add(person)

// 		// act
// 		row, err := r.GetById(1)
// 		if err != nil {
// 			t.Fatalf("expected nil, got %s", err)
// 		}

// 		entity := new(entity.Person)
// 		if err := row.Scan(entity.Id, entity.Name); err != nil {
// 			t.Fatalf("expected nil, got %s", err)
// 		}

// 		// assert
// 		if entity.GetId() != 1 {
// 			t.Fatalf("expected 1, got %d", entity.GetId())
// 		}
// 		if entity.Name != PERSON_NAME {
// 			t.Fatalf("expected Person 1, got %s", entity.Name)
// 		}
// 	})

// 	t.Run("search bigger entity", func(t *testing.T) {
// 		// arrange
// 		u := entity.User{
// 			Name:  "Person 1",
// 			Email: "123",
// 		}
// 		ru.Add(u)

// 		// act
// 		_, err := r.GetById(1)

// 		// assert
// 		if err != nil {
// 			t.Fatalf("expected nil, got %s", err)
// 		}
// 	})

// 	t.Run("not found", func(t *testing.T) {
// 		// arrange
// 		person := entity.Person{
// 			Name: "Person 1",
// 		}
// 		r.Add(person)

// 		// act
// 		res, _ := r.GetById(2)

// 		// assert
// 		if res != nil {
// 			t.Fatalf("expected error, got nil")
// 		}
// 	})
// }

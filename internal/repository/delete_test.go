package repository_test

// import (
// 	"testing"
// 	"time"

// 	"github.com/andresmeireles/speaker/internal/db"
// 	"github.com/andresmeireles/speaker/internal/db/entity"
// 	"github.com/andresmeireles/speaker/internal/repository"
// )

// func TestDelete(t *testing.T) {
// 	d, _ := db.GetDB()
// 	r := repository.Repository[entity.Person]{
// 		Db: d,
// 	}
// 	ri := repository.Repository[entity.Invite]{
// 		Db: d,
// 	}
// 	ru := repository.Repository[entity.User]{
// 		Db: d,
// 	}

// 	t.Run("delete register", func(t *testing.T) {
// 		// arrange
// 		user := entity.User{
// 			Name:  "Person 1",
// 			Email: "123",
// 		}
// 		ru.Add(user)

// 		// act
// 		err := ru.Delete(user)
// 		allUsers, gErr := r.GetAll()

// 		// assert
// 		if err != nil {
// 			t.Fatalf("expected nil, got %s", err)
// 		}

// 		if gErr != nil {
// 			t.Fatalf("expected nil, got %s", gErr)
// 		}

// 		numberOfRegisters := 0
// 		for allUsers.Next() {
// 			numberOfRegisters++
// 		}

// 		if numberOfRegisters != 0 {
// 			t.Fatalf("expected 0, got %d", numberOfRegisters)
// 		}
// 	})

// 	t.Run("Test Delete With Relationships", func(t *testing.T) {
// 		// arrange
// 		person := entity.Person{
// 			Name: "Person 1",
// 		}
// 		invite := entity.Invite{
// 			PersonId:   1,
// 			Theme:      "Theme",
// 			Time:       1,
// 			Date:       time.Now(),
// 			Accepted:   true,
// 			Remembered: true,
// 		}
// 		r.Add(person)
// 		ri.Add(invite)

// 		// act
// 		err := r.Delete(person)

// 		// assert
// 		if err != nil {
// 			t.Fatalf("expected nil, got %s", err)
// 		}
// 	})
// }

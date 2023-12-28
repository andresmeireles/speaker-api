package repository_test

import (
	"testing"

	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/db/repository"
	"github.com/andresmeireles/speaker/testdata"
)

func TestMain(m *testing.M) {
	testdata.SetupDatabase(m)
}

func TestAdd(t *testing.T) {
	somePerson := entity.Person{
		Name: "Some Person",
	}

	err := repository.Add(somePerson)

	if err != nil {
		t.Fatalf("expected nil, got %s", err)
	}
}

// func TestAddWithRelationships(t *testing.T) {
// 	// arrange
// 	person := entity.Person{
// 		Name: "Some Person",
// 	}
// 	err := repository.Add(person)

// 	if err != nil {
// 		t.Fatalf("error when create person, expected nil, got %s", err)
// 	}

// 	dbPerson := repository.GetById[entity.Person](1)

// 	if dbPerson == nil {
// 		t.Fatalf("expected not nil, got nil")
// 	}

// 	p := new(entity.Person)

// 	if err := dbPerson.Scan(&p.Id, &p.Name); err != nil {
// 		t.Fatalf("error when scan person, expected nil, got %s", err)
// 	}

// 	if err != nil {
// 		t.Fatalf("error when get person, expected nil, got %s", err)
// 	}

// 	invite := entity.Invite{
// 		Person:     *p,
// 		Theme:      "Theme",
// 		Time:       1,
// 		Date:       time.Now(),
// 		Accepted:   true,
// 		Remembered: true,
// 	}

// 	// act
// 	err = repository.Add(invite)

// 	if err != nil {
// 		t.Fatalf("expected nil, got %s", err)
// 	}

// 	dbInvite := repository.GetById[entity.Invite](1)

// 	if dbInvite == nil {
// 		t.Fatalf("expected not nil, got nil")
// 	}

// 	i := new(entity.Invite)

// 	if err := dbInvite.Scan(
// 		&i.Id,
// 		&i.Theme,
// 		&i.Person,
// 		i.Time,
// 		i.Date,
// 		i.Accepted,
// 		i.Remembered,
// 	); err != nil {
// 		t.Fatalf("error when scan invite, expected nil, got %s", err)
// 	}

// 	if err != nil {
// 		t.Fatalf("expected nil, got %s", err)
// 	}

// 	// assert
// 	if dbInvite == nil {
// 		t.Fatalf("expected not nil, got nil")
// 	}

// 	if i.Theme != "Theme" {
// 		t.Fatalf("expected Theme, got %s", i.Theme)
// 	}

// 	if i.Person.Name != "Some Person" {
// 		t.Fatalf("expected Some Person, got %s", i.Person.Name)
// 	}

// 	if i.Person.GetId() != 1 {
// 		t.Fatalf("expected 1, got %d", i.Person.GetId())
// 	}
// }

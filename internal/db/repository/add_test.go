package repository_test

import (
	"testing"
	"time"

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

func TestAddWithRelationships(t *testing.T) {
	// arrange
	person := entity.Person{
		Name: "Some Person",
	}
	_ = repository.Add(person)
	dbPerson, _ := repository.GetById[entity.Person](1, person)
	invite := entity.Invite{
		Person:     *dbPerson,
		Theme:      "Theme",
		Time:       1,
		Date:       int(time.Now().Unix()),
		Accepted:   true,
		Remembered: true,
	}

	// act
	_ = repository.Add(invite)
	dbInvite, _ := repository.GetById[entity.Invite](1, invite)

	// assert
	if dbInvite == nil {
		t.Fatalf("expected not nil, got nil")
	}
	if dbInvite.Person.GetId() != 1 {
		t.Fatalf("expected 1, got %d", dbInvite.Person.GetId())
	}
	if dbInvite.Person.Name != "Some Person" {
		t.Fatalf("expected Some Person, got %s", dbInvite.Person.Name)
	}
	if dbInvite.Theme != "Theme" {
		t.Fatalf("expected Theme, got %s", dbInvite.Theme)
	}
}

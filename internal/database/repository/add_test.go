package repository_test

import (
	"testing"

	"github.com/andresmeireles/speaker/internal/database/entity"
	"github.com/andresmeireles/speaker/internal/database/repository"
	"github.com/andresmeireles/speaker/testdata"
)

func TestMain(m *testing.M) {
	testdata.Setup(m)
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

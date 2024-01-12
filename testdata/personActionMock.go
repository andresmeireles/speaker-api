package testdata

import (
	"github.com/andresmeireles/speaker/internal/person"
	"github.com/stretchr/testify/mock"
)

type PersonActionMock struct {
	mock.Mock
}

func (a *PersonActionMock) Write(person person.Person) error {
	args := a.Called(person)
	return args.Error(0)
}

func (a *PersonActionMock) RemovePerson(personId int) error {
	args := a.Called(personId)
	return args.Error(0)
}

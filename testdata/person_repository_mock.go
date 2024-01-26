package testdata

import (
	"github.com/andresmeireles/speaker/internal/person"
	"github.com/andresmeireles/speaker/internal/tools/servicelocator"
	"github.com/stretchr/testify/mock"
)

type PersonRepositoryMock struct {
	mock.Mock
}

func (m *PersonRepositoryMock) New(s servicelocator.SL) any {
	args := m.Called(s)
	return args.Get(0)
}

func (m *PersonRepositoryMock) Add(person person.Person) error {
	args := m.Called(person)
	return args.Error(0)
}

func (m *PersonRepositoryMock) GetById(id int) (*person.Person, error) {
	args := m.Called(id)
	return args.Get(0).(*person.Person), args.Error(1)
}

func (m *PersonRepositoryMock) GetByName(name string) (*person.Person, error) {
	args := m.Called(name)
	return args.Get(0).(*person.Person), args.Error(1)
}

func (m *PersonRepositoryMock) Update(person person.Person) error {
	args := m.Called(person)
	return args.Error(0)
}

func (m *PersonRepositoryMock) GetAll() ([]person.Person, error) {
	args := m.Called()
	return args.Get(0).([]person.Person), args.Error(1)
}

func (m *PersonRepositoryMock) Delete(person person.Person) error {
	args := m.Called(person)
	return args.Error(0)
}

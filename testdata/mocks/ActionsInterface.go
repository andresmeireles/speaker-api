// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	person "github.com/andresmeireles/speaker/internal/person"
	mock "github.com/stretchr/testify/mock"
)

// ActionsInterface is an autogenerated mock type for the ActionsInterface type
type ActionsInterface struct {
	mock.Mock
}

type ActionsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *ActionsInterface) EXPECT() *ActionsInterface_Expecter {
	return &ActionsInterface_Expecter{mock: &_m.Mock}
}

// RemovePerson provides a mock function with given fields: personId
func (_m *ActionsInterface) RemovePerson(personId int) error {
	ret := _m.Called(personId)

	if len(ret) == 0 {
		panic("no return value specified for RemovePerson")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(personId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ActionsInterface_RemovePerson_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemovePerson'
type ActionsInterface_RemovePerson_Call struct {
	*mock.Call
}

// RemovePerson is a helper method to define mock.On call
//   - personId int
func (_e *ActionsInterface_Expecter) RemovePerson(personId interface{}) *ActionsInterface_RemovePerson_Call {
	return &ActionsInterface_RemovePerson_Call{Call: _e.mock.On("RemovePerson", personId)}
}

func (_c *ActionsInterface_RemovePerson_Call) Run(run func(personId int)) *ActionsInterface_RemovePerson_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *ActionsInterface_RemovePerson_Call) Return(_a0 error) *ActionsInterface_RemovePerson_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ActionsInterface_RemovePerson_Call) RunAndReturn(run func(int) error) *ActionsInterface_RemovePerson_Call {
	_c.Call.Return(run)
	return _c
}

// Write provides a mock function with given fields: _a0
func (_m *ActionsInterface) Write(_a0 person.Person) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Write")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(person.Person) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ActionsInterface_Write_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Write'
type ActionsInterface_Write_Call struct {
	*mock.Call
}

// Write is a helper method to define mock.On call
//   - _a0 person.Person
func (_e *ActionsInterface_Expecter) Write(_a0 interface{}) *ActionsInterface_Write_Call {
	return &ActionsInterface_Write_Call{Call: _e.mock.On("Write", _a0)}
}

func (_c *ActionsInterface_Write_Call) Run(run func(_a0 person.Person)) *ActionsInterface_Write_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(person.Person))
	})
	return _c
}

func (_c *ActionsInterface_Write_Call) Return(_a0 error) *ActionsInterface_Write_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ActionsInterface_Write_Call) RunAndReturn(run func(person.Person) error) *ActionsInterface_Write_Call {
	_c.Call.Return(run)
	return _c
}

// NewActionsInterface creates a new instance of ActionsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewActionsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ActionsInterface {
	mock := &ActionsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	user "github.com/andresmeireles/speaker/internal/user"
	mock "github.com/stretchr/testify/mock"
)

// UserRepositoryuser is an autogenerated mock type for the UserRepository type
type UserRepositoryuser struct {
	mock.Mock
}

type UserRepositoryuser_Expecter struct {
	mock *mock.Mock
}

func (_m *UserRepositoryuser) EXPECT() *UserRepositoryuser_Expecter {
	return &UserRepositoryuser_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: _a0
func (_m *UserRepositoryuser) Add(_a0 user.User) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(user.User) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserRepositoryuser_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type UserRepositoryuser_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - _a0 user.User
func (_e *UserRepositoryuser_Expecter) Add(_a0 interface{}) *UserRepositoryuser_Add_Call {
	return &UserRepositoryuser_Add_Call{Call: _e.mock.On("Add", _a0)}
}

func (_c *UserRepositoryuser_Add_Call) Run(run func(_a0 user.User)) *UserRepositoryuser_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(user.User))
	})
	return _c
}

func (_c *UserRepositoryuser_Add_Call) Return(_a0 error) *UserRepositoryuser_Add_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UserRepositoryuser_Add_Call) RunAndReturn(run func(user.User) error) *UserRepositoryuser_Add_Call {
	_c.Call.Return(run)
	return _c
}

// GetByEmail provides a mock function with given fields: email
func (_m *UserRepositoryuser) GetByEmail(email string) (user.User, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for GetByEmail")
	}

	var r0 user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (user.User, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) user.User); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(user.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepositoryuser_GetByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByEmail'
type UserRepositoryuser_GetByEmail_Call struct {
	*mock.Call
}

// GetByEmail is a helper method to define mock.On call
//   - email string
func (_e *UserRepositoryuser_Expecter) GetByEmail(email interface{}) *UserRepositoryuser_GetByEmail_Call {
	return &UserRepositoryuser_GetByEmail_Call{Call: _e.mock.On("GetByEmail", email)}
}

func (_c *UserRepositoryuser_GetByEmail_Call) Run(run func(email string)) *UserRepositoryuser_GetByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *UserRepositoryuser_GetByEmail_Call) Return(_a0 user.User, _a1 error) *UserRepositoryuser_GetByEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserRepositoryuser_GetByEmail_Call) RunAndReturn(run func(string) (user.User, error)) *UserRepositoryuser_GetByEmail_Call {
	_c.Call.Return(run)
	return _c
}

// GetById provides a mock function with given fields: id
func (_m *UserRepositoryuser) GetById(id int) (user.User, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (user.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) user.User); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(user.User)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepositoryuser_GetById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetById'
type UserRepositoryuser_GetById_Call struct {
	*mock.Call
}

// GetById is a helper method to define mock.On call
//   - id int
func (_e *UserRepositoryuser_Expecter) GetById(id interface{}) *UserRepositoryuser_GetById_Call {
	return &UserRepositoryuser_GetById_Call{Call: _e.mock.On("GetById", id)}
}

func (_c *UserRepositoryuser_GetById_Call) Run(run func(id int)) *UserRepositoryuser_GetById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *UserRepositoryuser_GetById_Call) Return(_a0 user.User, _a1 error) *UserRepositoryuser_GetById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserRepositoryuser_GetById_Call) RunAndReturn(run func(int) (user.User, error)) *UserRepositoryuser_GetById_Call {
	_c.Call.Return(run)
	return _c
}

// NewUserRepositoryuser creates a new instance of UserRepositoryuser. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepositoryuser(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepositoryuser {
	mock := &UserRepositoryuser{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

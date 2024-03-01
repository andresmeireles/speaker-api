// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	auth "github.com/andresmeireles/speaker/internal/auth"
	mock "github.com/stretchr/testify/mock"

	user "github.com/andresmeireles/speaker/internal/user"
)

// Serviceauth is an autogenerated mock type for the Service type
type Serviceauth struct {
	mock.Mock
}

type Serviceauth_Expecter struct {
	mock *mock.Mock
}

func (_m *Serviceauth) EXPECT() *Serviceauth_Expecter {
	return &Serviceauth_Expecter{mock: &_m.Mock}
}

// CreateJWT provides a mock function with given fields: _a0, remember
func (_m *Serviceauth) CreateJWT(_a0 user.User, remember bool) (auth.Auth, error) {
	ret := _m.Called(_a0, remember)

	if len(ret) == 0 {
		panic("no return value specified for CreateJWT")
	}

	var r0 auth.Auth
	var r1 error
	if rf, ok := ret.Get(0).(func(user.User, bool) (auth.Auth, error)); ok {
		return rf(_a0, remember)
	}
	if rf, ok := ret.Get(0).(func(user.User, bool) auth.Auth); ok {
		r0 = rf(_a0, remember)
	} else {
		r0 = ret.Get(0).(auth.Auth)
	}

	if rf, ok := ret.Get(1).(func(user.User, bool) error); ok {
		r1 = rf(_a0, remember)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Serviceauth_CreateJWT_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateJWT'
type Serviceauth_CreateJWT_Call struct {
	*mock.Call
}

// CreateJWT is a helper method to define mock.On call
//   - _a0 user.User
//   - remember bool
func (_e *Serviceauth_Expecter) CreateJWT(_a0 interface{}, remember interface{}) *Serviceauth_CreateJWT_Call {
	return &Serviceauth_CreateJWT_Call{Call: _e.mock.On("CreateJWT", _a0, remember)}
}

func (_c *Serviceauth_CreateJWT_Call) Run(run func(_a0 user.User, remember bool)) *Serviceauth_CreateJWT_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(user.User), args[1].(bool))
	})
	return _c
}

func (_c *Serviceauth_CreateJWT_Call) Return(_a0 auth.Auth, _a1 error) *Serviceauth_CreateJWT_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Serviceauth_CreateJWT_Call) RunAndReturn(run func(user.User, bool) (auth.Auth, error)) *Serviceauth_CreateJWT_Call {
	_c.Call.Return(run)
	return _c
}

// HasEmail provides a mock function with given fields: email
func (_m *Serviceauth) HasEmail(email string) bool {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for HasEmail")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Serviceauth_HasEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HasEmail'
type Serviceauth_HasEmail_Call struct {
	*mock.Call
}

// HasEmail is a helper method to define mock.On call
//   - email string
func (_e *Serviceauth_Expecter) HasEmail(email interface{}) *Serviceauth_HasEmail_Call {
	return &Serviceauth_HasEmail_Call{Call: _e.mock.On("HasEmail", email)}
}

func (_c *Serviceauth_HasEmail_Call) Run(run func(email string)) *Serviceauth_HasEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Serviceauth_HasEmail_Call) Return(_a0 bool) *Serviceauth_HasEmail_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Serviceauth_HasEmail_Call) RunAndReturn(run func(string) bool) *Serviceauth_HasEmail_Call {
	_c.Call.Return(run)
	return _c
}

// Logout provides a mock function with given fields: userId
func (_m *Serviceauth) Logout(userId int) error {
	ret := _m.Called(userId)

	if len(ret) == 0 {
		panic("no return value specified for Logout")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Serviceauth_Logout_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Logout'
type Serviceauth_Logout_Call struct {
	*mock.Call
}

// Logout is a helper method to define mock.On call
//   - userId int
func (_e *Serviceauth_Expecter) Logout(userId interface{}) *Serviceauth_Logout_Call {
	return &Serviceauth_Logout_Call{Call: _e.mock.On("Logout", userId)}
}

func (_c *Serviceauth_Logout_Call) Run(run func(userId int)) *Serviceauth_Logout_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *Serviceauth_Logout_Call) Return(_a0 error) *Serviceauth_Logout_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Serviceauth_Logout_Call) RunAndReturn(run func(int) error) *Serviceauth_Logout_Call {
	_c.Call.Return(run)
	return _c
}

// SendCode provides a mock function with given fields: email
func (_m *Serviceauth) SendCode(email string) error {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for SendCode")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Serviceauth_SendCode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendCode'
type Serviceauth_SendCode_Call struct {
	*mock.Call
}

// SendCode is a helper method to define mock.On call
//   - email string
func (_e *Serviceauth_Expecter) SendCode(email interface{}) *Serviceauth_SendCode_Call {
	return &Serviceauth_SendCode_Call{Call: _e.mock.On("SendCode", email)}
}

func (_c *Serviceauth_SendCode_Call) Run(run func(email string)) *Serviceauth_SendCode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Serviceauth_SendCode_Call) Return(_a0 error) *Serviceauth_SendCode_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Serviceauth_SendCode_Call) RunAndReturn(run func(string) error) *Serviceauth_SendCode_Call {
	_c.Call.Return(run)
	return _c
}

// ValidateJwt provides a mock function with given fields: token
func (_m *Serviceauth) ValidateJwt(token string) bool {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for ValidateJwt")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Serviceauth_ValidateJwt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ValidateJwt'
type Serviceauth_ValidateJwt_Call struct {
	*mock.Call
}

// ValidateJwt is a helper method to define mock.On call
//   - token string
func (_e *Serviceauth_Expecter) ValidateJwt(token interface{}) *Serviceauth_ValidateJwt_Call {
	return &Serviceauth_ValidateJwt_Call{Call: _e.mock.On("ValidateJwt", token)}
}

func (_c *Serviceauth_ValidateJwt_Call) Run(run func(token string)) *Serviceauth_ValidateJwt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Serviceauth_ValidateJwt_Call) Return(_a0 bool) *Serviceauth_ValidateJwt_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Serviceauth_ValidateJwt_Call) RunAndReturn(run func(string) bool) *Serviceauth_ValidateJwt_Call {
	_c.Call.Return(run)
	return _c
}

// NewServiceauth creates a new instance of Serviceauth. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServiceauth(t interface {
	mock.TestingT
	Cleanup(func())
}) *Serviceauth {
	mock := &Serviceauth{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

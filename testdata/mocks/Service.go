// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	invite "github.com/andresmeireles/speaker/internal/invite"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

type Service_Expecter struct {
	mock *mock.Mock
}

func (_m *Service) EXPECT() *Service_Expecter {
	return &Service_Expecter{mock: &_m.Mock}
}

// AcceptInvite provides a mock function with given fields: inviteId
func (_m *Service) AcceptInvite(inviteId int) error {
	ret := _m.Called(inviteId)

	if len(ret) == 0 {
		panic("no return value specified for AcceptInvite")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(inviteId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Service_AcceptInvite_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AcceptInvite'
type Service_AcceptInvite_Call struct {
	*mock.Call
}

// AcceptInvite is a helper method to define mock.On call
//   - inviteId int
func (_e *Service_Expecter) AcceptInvite(inviteId interface{}) *Service_AcceptInvite_Call {
	return &Service_AcceptInvite_Call{Call: _e.mock.On("AcceptInvite", inviteId)}
}

func (_c *Service_AcceptInvite_Call) Run(run func(inviteId int)) *Service_AcceptInvite_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *Service_AcceptInvite_Call) Return(_a0 error) *Service_AcceptInvite_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Service_AcceptInvite_Call) RunAndReturn(run func(int) error) *Service_AcceptInvite_Call {
	_c.Call.Return(run)
	return _c
}

// CreateInvite provides a mock function with given fields: inviteData
func (_m *Service) CreateInvite(inviteData invite.InvitePost) (invite.Invite, error) {
	ret := _m.Called(inviteData)

	if len(ret) == 0 {
		panic("no return value specified for CreateInvite")
	}

	var r0 invite.Invite
	var r1 error
	if rf, ok := ret.Get(0).(func(invite.InvitePost) (invite.Invite, error)); ok {
		return rf(inviteData)
	}
	if rf, ok := ret.Get(0).(func(invite.InvitePost) invite.Invite); ok {
		r0 = rf(inviteData)
	} else {
		r0 = ret.Get(0).(invite.Invite)
	}

	if rf, ok := ret.Get(1).(func(invite.InvitePost) error); ok {
		r1 = rf(inviteData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Service_CreateInvite_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateInvite'
type Service_CreateInvite_Call struct {
	*mock.Call
}

// CreateInvite is a helper method to define mock.On call
//   - inviteData invite.InvitePost
func (_e *Service_Expecter) CreateInvite(inviteData interface{}) *Service_CreateInvite_Call {
	return &Service_CreateInvite_Call{Call: _e.mock.On("CreateInvite", inviteData)}
}

func (_c *Service_CreateInvite_Call) Run(run func(inviteData invite.InvitePost)) *Service_CreateInvite_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(invite.InvitePost))
	})
	return _c
}

func (_c *Service_CreateInvite_Call) Return(_a0 invite.Invite, _a1 error) *Service_CreateInvite_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Service_CreateInvite_Call) RunAndReturn(run func(invite.InvitePost) (invite.Invite, error)) *Service_CreateInvite_Call {
	_c.Call.Return(run)
	return _c
}

// ParseInviteWithTemplate provides a mock function with given fields: inviteId
func (_m *Service) ParseInviteWithTemplate(inviteId int) (string, error) {
	ret := _m.Called(inviteId)

	if len(ret) == 0 {
		panic("no return value specified for ParseInviteWithTemplate")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (string, error)); ok {
		return rf(inviteId)
	}
	if rf, ok := ret.Get(0).(func(int) string); ok {
		r0 = rf(inviteId)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(inviteId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Service_ParseInviteWithTemplate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ParseInviteWithTemplate'
type Service_ParseInviteWithTemplate_Call struct {
	*mock.Call
}

// ParseInviteWithTemplate is a helper method to define mock.On call
//   - inviteId int
func (_e *Service_Expecter) ParseInviteWithTemplate(inviteId interface{}) *Service_ParseInviteWithTemplate_Call {
	return &Service_ParseInviteWithTemplate_Call{Call: _e.mock.On("ParseInviteWithTemplate", inviteId)}
}

func (_c *Service_ParseInviteWithTemplate_Call) Run(run func(inviteId int)) *Service_ParseInviteWithTemplate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *Service_ParseInviteWithTemplate_Call) Return(_a0 string, _a1 error) *Service_ParseInviteWithTemplate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Service_ParseInviteWithTemplate_Call) RunAndReturn(run func(int) (string, error)) *Service_ParseInviteWithTemplate_Call {
	_c.Call.Return(run)
	return _c
}

// ParseRememberMessage provides a mock function with given fields: inviteId
func (_m *Service) ParseRememberMessage(inviteId int) (string, error) {
	ret := _m.Called(inviteId)

	if len(ret) == 0 {
		panic("no return value specified for ParseRememberMessage")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (string, error)); ok {
		return rf(inviteId)
	}
	if rf, ok := ret.Get(0).(func(int) string); ok {
		r0 = rf(inviteId)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(inviteId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Service_ParseRememberMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ParseRememberMessage'
type Service_ParseRememberMessage_Call struct {
	*mock.Call
}

// ParseRememberMessage is a helper method to define mock.On call
//   - inviteId int
func (_e *Service_Expecter) ParseRememberMessage(inviteId interface{}) *Service_ParseRememberMessage_Call {
	return &Service_ParseRememberMessage_Call{Call: _e.mock.On("ParseRememberMessage", inviteId)}
}

func (_c *Service_ParseRememberMessage_Call) Run(run func(inviteId int)) *Service_ParseRememberMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *Service_ParseRememberMessage_Call) Return(_a0 string, _a1 error) *Service_ParseRememberMessage_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Service_ParseRememberMessage_Call) RunAndReturn(run func(int) (string, error)) *Service_ParseRememberMessage_Call {
	_c.Call.Return(run)
	return _c
}

// RememberInvite provides a mock function with given fields: inviteId
func (_m *Service) RememberInvite(inviteId int) error {
	ret := _m.Called(inviteId)

	if len(ret) == 0 {
		panic("no return value specified for RememberInvite")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(inviteId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Service_RememberInvite_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RememberInvite'
type Service_RememberInvite_Call struct {
	*mock.Call
}

// RememberInvite is a helper method to define mock.On call
//   - inviteId int
func (_e *Service_Expecter) RememberInvite(inviteId interface{}) *Service_RememberInvite_Call {
	return &Service_RememberInvite_Call{Call: _e.mock.On("RememberInvite", inviteId)}
}

func (_c *Service_RememberInvite_Call) Run(run func(inviteId int)) *Service_RememberInvite_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *Service_RememberInvite_Call) Return(_a0 error) *Service_RememberInvite_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Service_RememberInvite_Call) RunAndReturn(run func(int) error) *Service_RememberInvite_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveInvite provides a mock function with given fields: inviteId
func (_m *Service) RemoveInvite(inviteId int) error {
	ret := _m.Called(inviteId)

	if len(ret) == 0 {
		panic("no return value specified for RemoveInvite")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(inviteId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Service_RemoveInvite_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveInvite'
type Service_RemoveInvite_Call struct {
	*mock.Call
}

// RemoveInvite is a helper method to define mock.On call
//   - inviteId int
func (_e *Service_Expecter) RemoveInvite(inviteId interface{}) *Service_RemoveInvite_Call {
	return &Service_RemoveInvite_Call{Call: _e.mock.On("RemoveInvite", inviteId)}
}

func (_c *Service_RemoveInvite_Call) Run(run func(inviteId int)) *Service_RemoveInvite_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *Service_RemoveInvite_Call) Return(_a0 error) *Service_RemoveInvite_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Service_RemoveInvite_Call) RunAndReturn(run func(int) error) *Service_RemoveInvite_Call {
	_c.Call.Return(run)
	return _c
}

// SetDoneStatus provides a mock function with given fields: inviteId, done
func (_m *Service) SetDoneStatus(inviteId int, done bool) error {
	ret := _m.Called(inviteId, done)

	if len(ret) == 0 {
		panic("no return value specified for SetDoneStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int, bool) error); ok {
		r0 = rf(inviteId, done)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Service_SetDoneStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetDoneStatus'
type Service_SetDoneStatus_Call struct {
	*mock.Call
}

// SetDoneStatus is a helper method to define mock.On call
//   - inviteId int
//   - done bool
func (_e *Service_Expecter) SetDoneStatus(inviteId interface{}, done interface{}) *Service_SetDoneStatus_Call {
	return &Service_SetDoneStatus_Call{Call: _e.mock.On("SetDoneStatus", inviteId, done)}
}

func (_c *Service_SetDoneStatus_Call) Run(run func(inviteId int, done bool)) *Service_SetDoneStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(bool))
	})
	return _c
}

func (_c *Service_SetDoneStatus_Call) Return(_a0 error) *Service_SetDoneStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Service_SetDoneStatus_Call) RunAndReturn(run func(int, bool) error) *Service_SetDoneStatus_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateInvite provides a mock function with given fields: updateInviteData, inviteId
func (_m *Service) UpdateInvite(updateInviteData invite.UpdateInviteData, inviteId int) error {
	ret := _m.Called(updateInviteData, inviteId)

	if len(ret) == 0 {
		panic("no return value specified for UpdateInvite")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(invite.UpdateInviteData, int) error); ok {
		r0 = rf(updateInviteData, inviteId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Service_UpdateInvite_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateInvite'
type Service_UpdateInvite_Call struct {
	*mock.Call
}

// UpdateInvite is a helper method to define mock.On call
//   - updateInviteData invite.UpdateInviteData
//   - inviteId int
func (_e *Service_Expecter) UpdateInvite(updateInviteData interface{}, inviteId interface{}) *Service_UpdateInvite_Call {
	return &Service_UpdateInvite_Call{Call: _e.mock.On("UpdateInvite", updateInviteData, inviteId)}
}

func (_c *Service_UpdateInvite_Call) Run(run func(updateInviteData invite.UpdateInviteData, inviteId int)) *Service_UpdateInvite_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(invite.UpdateInviteData), args[1].(int))
	})
	return _c
}

func (_c *Service_UpdateInvite_Call) Return(_a0 error) *Service_UpdateInvite_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Service_UpdateInvite_Call) RunAndReturn(run func(invite.UpdateInviteData, int) error) *Service_UpdateInvite_Call {
	_c.Call.Return(run)
	return _c
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

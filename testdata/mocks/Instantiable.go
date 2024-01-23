// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	servicelocator "github.com/andresmeireles/speaker/internal/tools/servicelocator"
	mock "github.com/stretchr/testify/mock"
)

// Instantiable is an autogenerated mock type for the Instantiable type
type Instantiable struct {
	mock.Mock
}

type Instantiable_Expecter struct {
	mock *mock.Mock
}

func (_m *Instantiable) EXPECT() *Instantiable_Expecter {
	return &Instantiable_Expecter{mock: &_m.Mock}
}

// New provides a mock function with given fields: s
func (_m *Instantiable) New(s servicelocator.SL) interface{} {
	ret := _m.Called(s)

	if len(ret) == 0 {
		panic("no return value specified for New")
	}

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(servicelocator.SL) interface{}); ok {
		r0 = rf(s)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// Instantiable_New_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'New'
type Instantiable_New_Call struct {
	*mock.Call
}

// New is a helper method to define mock.On call
//   - s servicelocator.ServiceLocator
func (_e *Instantiable_Expecter) New(s interface{}) *Instantiable_New_Call {
	return &Instantiable_New_Call{Call: _e.mock.On("New", s)}
}

func (_c *Instantiable_New_Call) Run(run func(s servicelocator.SL)) *Instantiable_New_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(servicelocator.SL))
	})
	return _c
}

func (_c *Instantiable_New_Call) Return(_a0 interface{}) *Instantiable_New_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Instantiable_New_Call) RunAndReturn(run func(servicelocator.SL) interface{}) *Instantiable_New_Call {
	_c.Call.Return(run)
	return _c
}

// NewInstantiable creates a new instance of Instantiable. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewInstantiable(t interface {
	mock.TestingT
	Cleanup(func())
}) *Instantiable {
	mock := &Instantiable{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

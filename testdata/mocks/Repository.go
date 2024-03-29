// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	config "github.com/andresmeireles/speaker/internal/config"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

type Repository_Expecter struct {
	mock *mock.Mock
}

func (_m *Repository) EXPECT() *Repository_Expecter {
	return &Repository_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: _a0
func (_m *Repository) Add(_a0 config.Config) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(config.Config) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type Repository_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - _a0 config.Config
func (_e *Repository_Expecter) Add(_a0 interface{}) *Repository_Add_Call {
	return &Repository_Add_Call{Call: _e.mock.On("Add", _a0)}
}

func (_c *Repository_Add_Call) Run(run func(_a0 config.Config)) *Repository_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(config.Config))
	})
	return _c
}

func (_c *Repository_Add_Call) Return(_a0 error) *Repository_Add_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_Add_Call) RunAndReturn(run func(config.Config) error) *Repository_Add_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: _a0
func (_m *Repository) Delete(_a0 config.Config) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(config.Config) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type Repository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - _a0 config.Config
func (_e *Repository_Expecter) Delete(_a0 interface{}) *Repository_Delete_Call {
	return &Repository_Delete_Call{Call: _e.mock.On("Delete", _a0)}
}

func (_c *Repository_Delete_Call) Run(run func(_a0 config.Config)) *Repository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(config.Config))
	})
	return _c
}

func (_c *Repository_Delete_Call) Return(_a0 error) *Repository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_Delete_Call) RunAndReturn(run func(config.Config) error) *Repository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetAll provides a mock function with given fields:
func (_m *Repository) GetAll() ([]config.Config, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []config.Config
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]config.Config, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []config.Config); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]config.Config)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type Repository_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
func (_e *Repository_Expecter) GetAll() *Repository_GetAll_Call {
	return &Repository_GetAll_Call{Call: _e.mock.On("GetAll")}
}

func (_c *Repository_GetAll_Call) Run(run func()) *Repository_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Repository_GetAll_Call) Return(_a0 []config.Config, _a1 error) *Repository_GetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_GetAll_Call) RunAndReturn(run func() ([]config.Config, error)) *Repository_GetAll_Call {
	_c.Call.Return(run)
	return _c
}

// GetById provides a mock function with given fields: id
func (_m *Repository) GetById(id int) (*config.Config, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 *config.Config
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*config.Config, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *config.Config); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*config.Config)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_GetById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetById'
type Repository_GetById_Call struct {
	*mock.Call
}

// GetById is a helper method to define mock.On call
//   - id int
func (_e *Repository_Expecter) GetById(id interface{}) *Repository_GetById_Call {
	return &Repository_GetById_Call{Call: _e.mock.On("GetById", id)}
}

func (_c *Repository_GetById_Call) Run(run func(id int)) *Repository_GetById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *Repository_GetById_Call) Return(_a0 *config.Config, _a1 error) *Repository_GetById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_GetById_Call) RunAndReturn(run func(int) (*config.Config, error)) *Repository_GetById_Call {
	_c.Call.Return(run)
	return _c
}

// GetByName provides a mock function with given fields: name
func (_m *Repository) GetByName(name string) (*config.Config, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for GetByName")
	}

	var r0 *config.Config
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*config.Config, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) *config.Config); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*config.Config)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_GetByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByName'
type Repository_GetByName_Call struct {
	*mock.Call
}

// GetByName is a helper method to define mock.On call
//   - name string
func (_e *Repository_Expecter) GetByName(name interface{}) *Repository_GetByName_Call {
	return &Repository_GetByName_Call{Call: _e.mock.On("GetByName", name)}
}

func (_c *Repository_GetByName_Call) Run(run func(name string)) *Repository_GetByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Repository_GetByName_Call) Return(_a0 *config.Config, _a1 error) *Repository_GetByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_GetByName_Call) RunAndReturn(run func(string) (*config.Config, error)) *Repository_GetByName_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: _a0
func (_m *Repository) Update(_a0 config.Config) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(config.Config) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type Repository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - _a0 config.Config
func (_e *Repository_Expecter) Update(_a0 interface{}) *Repository_Update_Call {
	return &Repository_Update_Call{Call: _e.mock.On("Update", _a0)}
}

func (_c *Repository_Update_Call) Run(run func(_a0 config.Config)) *Repository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(config.Config))
	})
	return _c
}

func (_c *Repository_Update_Call) Return(_a0 error) *Repository_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_Update_Call) RunAndReturn(run func(config.Config) error) *Repository_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

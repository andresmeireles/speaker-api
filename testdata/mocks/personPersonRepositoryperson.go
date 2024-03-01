// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	person "github.com/andresmeireles/speaker/internal/person"
	mock "github.com/stretchr/testify/mock"
)

// PersonRepositoryperson is an autogenerated mock type for the PersonRepository type
type PersonRepositoryperson struct {
	mock.Mock
}

type PersonRepositoryperson_Expecter struct {
	mock *mock.Mock
}

func (_m *PersonRepositoryperson) EXPECT() *PersonRepositoryperson_Expecter {
	return &PersonRepositoryperson_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: _a0
func (_m *PersonRepositoryperson) Add(_a0 person.Person) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(person.Person) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PersonRepositoryperson_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type PersonRepositoryperson_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - _a0 person.Person
func (_e *PersonRepositoryperson_Expecter) Add(_a0 interface{}) *PersonRepositoryperson_Add_Call {
	return &PersonRepositoryperson_Add_Call{Call: _e.mock.On("Add", _a0)}
}

func (_c *PersonRepositoryperson_Add_Call) Run(run func(_a0 person.Person)) *PersonRepositoryperson_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(person.Person))
	})
	return _c
}

func (_c *PersonRepositoryperson_Add_Call) Return(_a0 error) *PersonRepositoryperson_Add_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PersonRepositoryperson_Add_Call) RunAndReturn(run func(person.Person) error) *PersonRepositoryperson_Add_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: _a0
func (_m *PersonRepositoryperson) Delete(_a0 person.Person) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(person.Person) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PersonRepositoryperson_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type PersonRepositoryperson_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - _a0 person.Person
func (_e *PersonRepositoryperson_Expecter) Delete(_a0 interface{}) *PersonRepositoryperson_Delete_Call {
	return &PersonRepositoryperson_Delete_Call{Call: _e.mock.On("Delete", _a0)}
}

func (_c *PersonRepositoryperson_Delete_Call) Run(run func(_a0 person.Person)) *PersonRepositoryperson_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(person.Person))
	})
	return _c
}

func (_c *PersonRepositoryperson_Delete_Call) Return(_a0 error) *PersonRepositoryperson_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PersonRepositoryperson_Delete_Call) RunAndReturn(run func(person.Person) error) *PersonRepositoryperson_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetAll provides a mock function with given fields:
func (_m *PersonRepositoryperson) GetAll() ([]person.Person, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []person.Person
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]person.Person, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []person.Person); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]person.Person)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PersonRepositoryperson_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type PersonRepositoryperson_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
func (_e *PersonRepositoryperson_Expecter) GetAll() *PersonRepositoryperson_GetAll_Call {
	return &PersonRepositoryperson_GetAll_Call{Call: _e.mock.On("GetAll")}
}

func (_c *PersonRepositoryperson_GetAll_Call) Run(run func()) *PersonRepositoryperson_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PersonRepositoryperson_GetAll_Call) Return(_a0 []person.Person, _a1 error) *PersonRepositoryperson_GetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PersonRepositoryperson_GetAll_Call) RunAndReturn(run func() ([]person.Person, error)) *PersonRepositoryperson_GetAll_Call {
	_c.Call.Return(run)
	return _c
}

// GetById provides a mock function with given fields: id
func (_m *PersonRepositoryperson) GetById(id int) (*person.Person, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 *person.Person
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*person.Person, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *person.Person); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*person.Person)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PersonRepositoryperson_GetById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetById'
type PersonRepositoryperson_GetById_Call struct {
	*mock.Call
}

// GetById is a helper method to define mock.On call
//   - id int
func (_e *PersonRepositoryperson_Expecter) GetById(id interface{}) *PersonRepositoryperson_GetById_Call {
	return &PersonRepositoryperson_GetById_Call{Call: _e.mock.On("GetById", id)}
}

func (_c *PersonRepositoryperson_GetById_Call) Run(run func(id int)) *PersonRepositoryperson_GetById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *PersonRepositoryperson_GetById_Call) Return(_a0 *person.Person, _a1 error) *PersonRepositoryperson_GetById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PersonRepositoryperson_GetById_Call) RunAndReturn(run func(int) (*person.Person, error)) *PersonRepositoryperson_GetById_Call {
	_c.Call.Return(run)
	return _c
}

// GetByName provides a mock function with given fields: name
func (_m *PersonRepositoryperson) GetByName(name string) (*person.Person, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for GetByName")
	}

	var r0 *person.Person
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*person.Person, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) *person.Person); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*person.Person)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PersonRepositoryperson_GetByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByName'
type PersonRepositoryperson_GetByName_Call struct {
	*mock.Call
}

// GetByName is a helper method to define mock.On call
//   - name string
func (_e *PersonRepositoryperson_Expecter) GetByName(name interface{}) *PersonRepositoryperson_GetByName_Call {
	return &PersonRepositoryperson_GetByName_Call{Call: _e.mock.On("GetByName", name)}
}

func (_c *PersonRepositoryperson_GetByName_Call) Run(run func(name string)) *PersonRepositoryperson_GetByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *PersonRepositoryperson_GetByName_Call) Return(_a0 *person.Person, _a1 error) *PersonRepositoryperson_GetByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PersonRepositoryperson_GetByName_Call) RunAndReturn(run func(string) (*person.Person, error)) *PersonRepositoryperson_GetByName_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: _a0
func (_m *PersonRepositoryperson) Update(_a0 person.Person) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(person.Person) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PersonRepositoryperson_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type PersonRepositoryperson_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - _a0 person.Person
func (_e *PersonRepositoryperson_Expecter) Update(_a0 interface{}) *PersonRepositoryperson_Update_Call {
	return &PersonRepositoryperson_Update_Call{Call: _e.mock.On("Update", _a0)}
}

func (_c *PersonRepositoryperson_Update_Call) Run(run func(_a0 person.Person)) *PersonRepositoryperson_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(person.Person))
	})
	return _c
}

func (_c *PersonRepositoryperson_Update_Call) Return(_a0 error) *PersonRepositoryperson_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PersonRepositoryperson_Update_Call) RunAndReturn(run func(person.Person) error) *PersonRepositoryperson_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewPersonRepositoryperson creates a new instance of PersonRepositoryperson. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPersonRepositoryperson(t interface {
	mock.TestingT
	Cleanup(func())
}) *PersonRepositoryperson {
	mock := &PersonRepositoryperson{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

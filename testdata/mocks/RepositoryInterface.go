// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	db "github.com/andresmeireles/speaker/internal/db"
	mock "github.com/stretchr/testify/mock"

	sql "database/sql"
)

// RepositoryInterface is an autogenerated mock type for the RepositoryInterface type
type RepositoryInterface struct {
	mock.Mock
}

type RepositoryInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *RepositoryInterface) EXPECT() *RepositoryInterface_Expecter {
	return &RepositoryInterface_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: en
func (_m *RepositoryInterface) Add(en db.Entity) error {
	ret := _m.Called(en)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(db.Entity) error); ok {
		r0 = rf(en)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RepositoryInterface_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type RepositoryInterface_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - en db.Entity
func (_e *RepositoryInterface_Expecter) Add(en interface{}) *RepositoryInterface_Add_Call {
	return &RepositoryInterface_Add_Call{Call: _e.mock.On("Add", en)}
}

func (_c *RepositoryInterface_Add_Call) Run(run func(en db.Entity)) *RepositoryInterface_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(db.Entity))
	})
	return _c
}

func (_c *RepositoryInterface_Add_Call) Return(_a0 error) *RepositoryInterface_Add_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RepositoryInterface_Add_Call) RunAndReturn(run func(db.Entity) error) *RepositoryInterface_Add_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: en
func (_m *RepositoryInterface) Delete(en db.Entity) error {
	ret := _m.Called(en)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(db.Entity) error); ok {
		r0 = rf(en)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RepositoryInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type RepositoryInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - en db.Entity
func (_e *RepositoryInterface_Expecter) Delete(en interface{}) *RepositoryInterface_Delete_Call {
	return &RepositoryInterface_Delete_Call{Call: _e.mock.On("Delete", en)}
}

func (_c *RepositoryInterface_Delete_Call) Run(run func(en db.Entity)) *RepositoryInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(db.Entity))
	})
	return _c
}

func (_c *RepositoryInterface_Delete_Call) Return(_a0 error) *RepositoryInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RepositoryInterface_Delete_Call) RunAndReturn(run func(db.Entity) error) *RepositoryInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetAll provides a mock function with given fields: table
func (_m *RepositoryInterface) GetAll(table string) (*sql.Rows, error) {
	ret := _m.Called(table)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 *sql.Rows
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*sql.Rows, error)); ok {
		return rf(table)
	}
	if rf, ok := ret.Get(0).(func(string) *sql.Rows); ok {
		r0 = rf(table)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Rows)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(table)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RepositoryInterface_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type RepositoryInterface_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
//   - table string
func (_e *RepositoryInterface_Expecter) GetAll(table interface{}) *RepositoryInterface_GetAll_Call {
	return &RepositoryInterface_GetAll_Call{Call: _e.mock.On("GetAll", table)}
}

func (_c *RepositoryInterface_GetAll_Call) Run(run func(table string)) *RepositoryInterface_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *RepositoryInterface_GetAll_Call) Return(_a0 *sql.Rows, _a1 error) *RepositoryInterface_GetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RepositoryInterface_GetAll_Call) RunAndReturn(run func(string) (*sql.Rows, error)) *RepositoryInterface_GetAll_Call {
	_c.Call.Return(run)
	return _c
}

// GetById provides a mock function with given fields: table, id
func (_m *RepositoryInterface) GetById(table string, id int) (*sql.Row, error) {
	ret := _m.Called(table, id)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 *sql.Row
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int) (*sql.Row, error)); ok {
		return rf(table, id)
	}
	if rf, ok := ret.Get(0).(func(string, int) *sql.Row); ok {
		r0 = rf(table, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Row)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(table, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RepositoryInterface_GetById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetById'
type RepositoryInterface_GetById_Call struct {
	*mock.Call
}

// GetById is a helper method to define mock.On call
//   - table string
//   - id int
func (_e *RepositoryInterface_Expecter) GetById(table interface{}, id interface{}) *RepositoryInterface_GetById_Call {
	return &RepositoryInterface_GetById_Call{Call: _e.mock.On("GetById", table, id)}
}

func (_c *RepositoryInterface_GetById_Call) Run(run func(table string, id int)) *RepositoryInterface_GetById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(int))
	})
	return _c
}

func (_c *RepositoryInterface_GetById_Call) Return(_a0 *sql.Row, _a1 error) *RepositoryInterface_GetById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RepositoryInterface_GetById_Call) RunAndReturn(run func(string, int) (*sql.Row, error)) *RepositoryInterface_GetById_Call {
	_c.Call.Return(run)
	return _c
}

// Query provides a mock function with given fields: q, args
func (_m *RepositoryInterface) Query(q string, args ...interface{}) (*sql.Rows, error) {
	var _ca []interface{}
	_ca = append(_ca, q)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Query")
	}

	var r0 *sql.Rows
	var r1 error
	if rf, ok := ret.Get(0).(func(string, ...interface{}) (*sql.Rows, error)); ok {
		return rf(q, args...)
	}
	if rf, ok := ret.Get(0).(func(string, ...interface{}) *sql.Rows); ok {
		r0 = rf(q, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Rows)
		}
	}

	if rf, ok := ret.Get(1).(func(string, ...interface{}) error); ok {
		r1 = rf(q, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RepositoryInterface_Query_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Query'
type RepositoryInterface_Query_Call struct {
	*mock.Call
}

// Query is a helper method to define mock.On call
//   - q string
//   - args ...interface{}
func (_e *RepositoryInterface_Expecter) Query(q interface{}, args ...interface{}) *RepositoryInterface_Query_Call {
	return &RepositoryInterface_Query_Call{Call: _e.mock.On("Query",
		append([]interface{}{q}, args...)...)}
}

func (_c *RepositoryInterface_Query_Call) Run(run func(q string, args ...interface{})) *RepositoryInterface_Query_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *RepositoryInterface_Query_Call) Return(_a0 *sql.Rows, _a1 error) *RepositoryInterface_Query_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RepositoryInterface_Query_Call) RunAndReturn(run func(string, ...interface{}) (*sql.Rows, error)) *RepositoryInterface_Query_Call {
	_c.Call.Return(run)
	return _c
}

// SingleQuery provides a mock function with given fields: q, args
func (_m *RepositoryInterface) SingleQuery(q string, args ...interface{}) (*sql.Row, error) {
	var _ca []interface{}
	_ca = append(_ca, q)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SingleQuery")
	}

	var r0 *sql.Row
	var r1 error
	if rf, ok := ret.Get(0).(func(string, ...interface{}) (*sql.Row, error)); ok {
		return rf(q, args...)
	}
	if rf, ok := ret.Get(0).(func(string, ...interface{}) *sql.Row); ok {
		r0 = rf(q, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Row)
		}
	}

	if rf, ok := ret.Get(1).(func(string, ...interface{}) error); ok {
		r1 = rf(q, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RepositoryInterface_SingleQuery_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SingleQuery'
type RepositoryInterface_SingleQuery_Call struct {
	*mock.Call
}

// SingleQuery is a helper method to define mock.On call
//   - q string
//   - args ...interface{}
func (_e *RepositoryInterface_Expecter) SingleQuery(q interface{}, args ...interface{}) *RepositoryInterface_SingleQuery_Call {
	return &RepositoryInterface_SingleQuery_Call{Call: _e.mock.On("SingleQuery",
		append([]interface{}{q}, args...)...)}
}

func (_c *RepositoryInterface_SingleQuery_Call) Run(run func(q string, args ...interface{})) *RepositoryInterface_SingleQuery_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *RepositoryInterface_SingleQuery_Call) Return(_a0 *sql.Row, _a1 error) *RepositoryInterface_SingleQuery_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RepositoryInterface_SingleQuery_Call) RunAndReturn(run func(string, ...interface{}) (*sql.Row, error)) *RepositoryInterface_SingleQuery_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: en
func (_m *RepositoryInterface) Update(en db.Entity) error {
	ret := _m.Called(en)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(db.Entity) error); ok {
		r0 = rf(en)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RepositoryInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type RepositoryInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - en db.Entity
func (_e *RepositoryInterface_Expecter) Update(en interface{}) *RepositoryInterface_Update_Call {
	return &RepositoryInterface_Update_Call{Call: _e.mock.On("Update", en)}
}

func (_c *RepositoryInterface_Update_Call) Run(run func(en db.Entity)) *RepositoryInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(db.Entity))
	})
	return _c
}

func (_c *RepositoryInterface_Update_Call) Return(_a0 error) *RepositoryInterface_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RepositoryInterface_Update_Call) RunAndReturn(run func(db.Entity) error) *RepositoryInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewRepositoryInterface creates a new instance of RepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepositoryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *RepositoryInterface {
	mock := &RepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

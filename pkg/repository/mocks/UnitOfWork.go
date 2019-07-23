// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import repository "github.com/lfourky/go-transaction-management/pkg/repository"

// UnitOfWork is an autogenerated mock type for the UnitOfWork type
type UnitOfWork struct {
	mock.Mock
}

// BeginTransaction provides a mock function with given fields:
func (_m *UnitOfWork) BeginTransaction() (repository.Store, error) {
	ret := _m.Called()

	var r0 repository.Store
	if rf, ok := ret.Get(0).(func() repository.Store); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(repository.Store)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Commit provides a mock function with given fields:
func (_m *UnitOfWork) Commit() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Rollback provides a mock function with given fields:
func (_m *UnitOfWork) Rollback() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

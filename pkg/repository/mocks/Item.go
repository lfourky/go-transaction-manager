// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/lfourky/go-transaction-management/pkg/model"

// Item is an autogenerated mock type for the Item type
type Item struct {
	mock.Mock
}

// Create provides a mock function with given fields: item
func (_m *Item) Create(item *model.Item) error {
	ret := _m.Called(item)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Item) error); ok {
		r0 = rf(item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: item
func (_m *Item) Delete(item *model.Item) error {
	ret := _m.Called(item)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Item) error); ok {
		r0 = rf(item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

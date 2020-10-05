// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/evt/simple-web-server/model"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// UserRepo is an autogenerated mock type for the UserRepo type
type UserRepo struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: _a0, _a1
func (_m *UserRepo) CreateUser(_a0 context.Context, _a1 *model.DBUser) (*model.DBUser, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *model.DBUser
	if rf, ok := ret.Get(0).(func(context.Context, *model.DBUser) *model.DBUser); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.DBUser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.DBUser) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUser provides a mock function with given fields: _a0, _a1
func (_m *UserRepo) DeleteUser(_a0 context.Context, _a1 uuid.UUID) (*model.DBUser, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *model.DBUser
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *model.DBUser); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.DBUser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: _a0, _a1
func (_m *UserRepo) GetUser(_a0 context.Context, _a1 uuid.UUID) (*model.DBUser, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *model.DBUser
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *model.DBUser); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.DBUser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: _a0, _a1
func (_m *UserRepo) UpdateUser(_a0 context.Context, _a1 *model.DBUser) (*model.DBUser, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *model.DBUser
	if rf, ok := ret.Get(0).(func(context.Context, *model.DBUser) *model.DBUser); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.DBUser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.DBUser) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

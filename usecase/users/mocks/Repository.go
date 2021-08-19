// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"
	users "injar/usecase/users"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetByID provides a mock function with given fields: ctx, ID
func (_m *Repository) GetByID(ctx context.Context, ID int) (users.Domain, error) {
	ret := _m.Called(ctx, ID)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int) users.Domain); ok {
		r0 = rf(ctx, ID)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUsername provides a mock function with given fields: ctx, username
func (_m *Repository) GetByUsername(ctx context.Context, username string) (users.Domain, error) {
	ret := _m.Called(ctx, username)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string) users.Domain); ok {
		r0 = rf(ctx, username)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: ctx, data
func (_m *Repository) Store(ctx context.Context, data *users.Domain) error {
	ret := _m.Called(ctx, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *users.Domain) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, usersDomain
func (_m *Repository) Update(ctx context.Context, usersDomain *users.Domain) (users.Domain, error) {
	ret := _m.Called(ctx, usersDomain)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *users.Domain) users.Domain); ok {
		r0 = rf(ctx, usersDomain)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *users.Domain) error); ok {
		r1 = rf(ctx, usersDomain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

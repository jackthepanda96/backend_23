// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	multipart "mime/multipart"

	mock "github.com/stretchr/testify/mock"

	users "apibe23/internal/features/users"
)

// Services is an autogenerated mock type for the Services type
type Services struct {
	mock.Mock
}

// Login provides a mock function with given fields: email, password
func (_m *Services) Login(email string, password string) (users.User, string, error) {
	ret := _m.Called(email, password)

	var r0 users.User
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string) (users.User, string, error)); ok {
		return rf(email, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) users.User); ok {
		r0 = rf(email, password)
	} else {
		r0 = ret.Get(0).(users.User)
	}

	if rf, ok := ret.Get(1).(func(string, string) string); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(email, password)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Register provides a mock function with given fields: newUser
func (_m *Services) Register(newUser users.User) error {
	ret := _m.Called(newUser)

	var r0 error
	if rf, ok := ret.Get(0).(func(users.User) error); ok {
		r0 = rf(newUser)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Register2 provides a mock function with given fields: newUser, file
func (_m *Services) Register2(newUser string, file *multipart.FileHeader) error {
	ret := _m.Called(newUser, file)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *multipart.FileHeader) error); ok {
		r0 = rf(newUser, file)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewServices creates a new instance of Services. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServices(t interface {
	mock.TestingT
	Cleanup(func())
}) *Services {
	mock := &Services{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

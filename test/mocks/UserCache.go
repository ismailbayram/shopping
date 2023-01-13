// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UserCache is an autogenerated mock type for the UserCache type
type UserCache struct {
	mock.Mock
}

// GetUserIDByVerificationToken provides a mock function with given fields: _a0
func (_m *UserCache) GetUserIDByVerificationToken(_a0 string) uint {
	ret := _m.Called(_a0)

	var r0 uint
	if rf, ok := ret.Get(0).(func(string) uint); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}

type mockConstructorTestingTNewUserCache interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserCache creates a new instance of UserCache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserCache(t mockConstructorTestingTNewUserCache) *UserCache {
	mock := &UserCache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

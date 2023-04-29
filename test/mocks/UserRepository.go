// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	models "github.com/ismailbayram/shopping/internal/users/models"
	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// All provides a mock function with given fields: _a0
func (_m *UserRepository) All(_a0 map[string]interface{}) ([]models.User, error) {
	ret := _m.Called(_a0)

	var r0 []models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(map[string]interface{}) ([]models.User, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(map[string]interface{}) []models.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(map[string]interface{}) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: _a0
func (_m *UserRepository) Create(_a0 models.User) (models.User, error) {
	ret := _m.Called(_a0)

	var r0 models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(models.User) (models.User, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(models.User) models.User); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	if rf, ok := ret.Get(1).(func(models.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByEmail provides a mock function with given fields: _a0
func (_m *UserRepository) GetByEmail(_a0 string) (models.User, error) {
	ret := _m.Called(_a0)

	var r0 models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (models.User, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) models.User); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: _a0
func (_m *UserRepository) GetByID(_a0 uint) (models.User, error) {
	ret := _m.Called(_a0)

	var r0 models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (models.User, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(uint) models.User); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByToken provides a mock function with given fields: _a0
func (_m *UserRepository) GetByToken(_a0 string) (models.User, error) {
	ret := _m.Called(_a0)

	var r0 models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (models.User, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) models.User); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0
func (_m *UserRepository) Update(_a0 models.User) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.User) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

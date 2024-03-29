// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ImageStorage is an autogenerated mock type for the ImageStorage type
type ImageStorage struct {
	mock.Mock
}

// Upload provides a mock function with given fields: _a0, _a1
func (_m *ImageStorage) Upload(_a0 string, _a1 []byte) (string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, []byte) (string, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(string, []byte) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, []byte) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Url provides a mock function with given fields: _a0
func (_m *ImageStorage) Url(_a0 string) string {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewImageStorage interface {
	mock.TestingT
	Cleanup(func())
}

// NewImageStorage creates a new instance of ImageStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewImageStorage(t mockConstructorTestingTNewImageStorage) *ImageStorage {
	mock := &ImageStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

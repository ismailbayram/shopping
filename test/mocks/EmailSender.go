// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// EmailSender is an autogenerated mock type for the EmailSender type
type EmailSender struct {
	mock.Mock
}

// SendWelcomeEmail provides a mock function with given fields: _a0
func (_m *EmailSender) SendWelcomeEmail(_a0 string) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewEmailSender interface {
	mock.TestingT
	Cleanup(func())
}

// NewEmailSender creates a new instance of EmailSender. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEmailSender(t mockConstructorTestingTNewEmailSender) *EmailSender {
	mock := &EmailSender{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

package application

import (
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockUserRepository struct {
	mock.Mock
}

func TestUserService_IsVerified(t *testing.T) {
	mockedUR := new(MockUserRepository)

	mockedUR.On("GetByID", 123).Return(true, nil)
	//us := NewUserService(mockedUR, mockedUR)
}

package application

import (
	"errors"
	domain "github.com/ismailbayram/shopping/pkg/users/domain/models"
	"github.com/ismailbayram/shopping/test/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserService_Register_Failed(t *testing.T) {
	mockedUR := &mocks.UserRepository{}
	US := NewUserService(mockedUR)

	mockedUR.On("GetByEmail", "iso@iso.com").Return(&domain.User{}, nil)
	err := US.Register("iso@iso.com", "123456", "ismail", "bayram")
	assert.True(t, errors.Is(err, domain.ErrorUserAlreadyExists))
}

func TestUserService_Register_Success(t *testing.T) {
	mockedUR := &mocks.UserRepository{}
	US := NewUserService(mockedUR)

	mockedUR.On("GetByEmail", "iso@iso.com").Return(nil, nil)
	mockedUR.On(
		"Create",
		&domain.User{
			Email:     "iso@iso.com",
			FirstName: "ismail",
			LastName:  "bayram",
			IsActive:  true,
			Password:  "123456",
		},
	).Return(
		&domain.User{
			ID:        3,
			Email:     "iso@iso.com",
			FirstName: "ismail",
			LastName:  "bayram",
			IsActive:  true,
			Password:  "123456",
		},
		nil,
	)
	err := US.Register("iso@iso.com", "123456", "ismail", "bayram")
	assert.Nil(t, err)
}

func TestUserService_GetByID(t *testing.T) {
	user := &domain.User{
		ID:       2,
		IsActive: true,
	}

	mockedUR := &mocks.UserRepository{}
	US := NewUserService(mockedUR)

	mockedUR.On("GetByID", uint(1)).Return(nil, domain.ErrorUserNotFound)
	userGot, err := US.GetByID(1)
	assert.Nil(t, userGot)
	assert.NotNil(t, err)

	mockedUR.On("GetByID", uint(2)).Return(user, nil)
	userGot, err = US.GetByID(2)
	assert.NotNil(t, userGot)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userGot.ID)
}

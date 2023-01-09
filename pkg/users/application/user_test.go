package application

import (
	"errors"
	domain "github.com/ismailbayram/shopping/pkg/users/domain/models"
	"github.com/ismailbayram/shopping/test/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	mockedUR = &mocks.UserRepository{}
	mockedER = &mocks.EmailRepository{}

	US = NewUserService(mockedUR, mockedER)

	userIsmail = &domain.User{
		ID:       2,
		IsActive: true,
	}
)

func TestUserService_GetByID(t *testing.T) {
	mockedUR.On("GetByID", uint(1)).Return(nil, errors.New("User Could Not Found"))
	user, err := US.GetByID(1)
	assert.Nil(t, user)
	assert.NotNil(t, err)

	mockedUR.On("GetByID", uint(2)).Return(userIsmail, nil)
	user, err = US.GetByID(2)
	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userIsmail.ID)
}

func TestUserService_IsVerified(t *testing.T) {
	mockedER.On("GetPrimaryOfUser", userIsmail).Return(nil, errors.New("Email Not Found"))
	isVerified := US.IsVerified(userIsmail)
	assert.False(t, isVerified)

	mockedER.On("GetPrimaryOfUser", userIsmail).Return(&domain.Email{IsVerified: false}, nil)
	isVerified = US.IsVerified(userIsmail)
	assert.False(t, isVerified)

	mockedER.On("GetPrimaryOfUser", userIsmail).Return(&domain.Email{IsVerified: true}, nil)
	isVerified = US.IsVerified(userIsmail)
	assert.False(t, isVerified)
}

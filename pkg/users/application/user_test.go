package application

import (
	"errors"
	domain "github.com/ismailbayram/shopping/pkg/users/domain/models"
	"github.com/ismailbayram/shopping/test/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserService_GetByID(t *testing.T) {
	user := &domain.User{
		ID:       2,
		IsActive: true,
	}

	mockedUR := &mocks.UserRepository{}
	mockedES := &mocks.EmailSender{}
	mockedUC := &mocks.UserCache{}
	US := NewUserService(mockedUR, mockedES, mockedUC)

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

func TestUserService_Register_Failed(t *testing.T) {
	mockedUR := &mocks.UserRepository{}
	mockedES := &mocks.EmailSender{}
	mockedUC := &mocks.UserCache{}
	US := NewUserService(mockedUR, mockedES, mockedUC)

	mockedUR.On("GetByEmail", "iso@iso.com").Return(&domain.User{}, nil)
	err := US.Register("iso@iso.com", "123456", "ismail", "bayram")
	assert.True(t, errors.Is(err, domain.ErrorUserAlreadyExists))
	// TODO: fix it
	//mockedES.AssertNotCalled(t, "SendWelcomeEmail", "iso@iso.com")
}

func TestUserService_Register_Success(t *testing.T) {
	mockedUR := &mocks.UserRepository{}
	mockedES := &mocks.EmailSender{}
	mockedUC := &mocks.UserCache{}
	US := NewUserService(mockedUR, mockedES, mockedUC)

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
	mockedES.On("SendWelcomeEmail", "iso@iso.com")
	err := US.Register("iso@iso.com", "123456", "ismail", "bayram")
	assert.Nil(t, err)
	mockedES.AssertCalled(t, "SendWelcomeEmail", "iso@iso.com")
}

func TestUserService_Verify_Failed(t *testing.T) {
	mockedUR := &mocks.UserRepository{}
	mockedES := &mocks.EmailSender{}
	mockedUC := &mocks.UserCache{}
	US := NewUserService(mockedUR, mockedES, mockedUC)

	mockedUC.On("GetUserIDByVerificationToken", "token").Return(nil)
	err := US.Verify("token")
	assert.Equal(t, domain.ErrorUserNotFound, err)

	//cachedID := uint(1)
	//mockedUC.On("GetUserIDByVerificationToken", "token").Return(&cachedID)
	//mockedUR.On("GetByID", uint(1)).Return(nil, errors.New("dummy"))
	//err = US.Verify("token")
	//assert.Equal(t, domain.ErrorGeneral, err)

}

func TestUserService_Verify_Success(t *testing.T) {
	mockedUR := &mocks.UserRepository{}
	mockedES := &mocks.EmailSender{}
	mockedUC := &mocks.UserCache{}
	US := NewUserService(mockedUR, mockedES, mockedUC)

	cachedID := uint(1)
	mockedUC.On("GetUserIDByVerificationToken", "token").Return(&cachedID)
	mockedUR.On("GetByID", uint(1)).Return(&domain.User{ID: cachedID}, nil)
	mockedUR.On("Update", &domain.User{ID: cachedID, IsVerified: true}).Return(nil)
	err := US.Verify("token")
	assert.Nil(t, err)
	mockedUR.AssertCalled(t, "Update", &domain.User{ID: cachedID, IsVerified: true})

}

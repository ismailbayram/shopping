package application

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	domain "github.com/ismailbayram/shopping/pkg/users/domain/models"
	"github.com/ismailbayram/shopping/test/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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

func TestUserService_GetByToken(t *testing.T) {
	user := &domain.User{
		ID:       2,
		IsActive: true,
	}

	mockedUR := &mocks.UserRepository{}
	mockedES := &mocks.EmailSender{}
	mockedUC := &mocks.UserCache{}
	US := NewUserService(mockedUR, mockedES, mockedUC)

	mockedUR.On("GetByToken", "token1").Return(nil, domain.ErrorUserNotFound)
	userGot, err := US.GetByToken("token1")
	assert.Nil(t, userGot)
	assert.NotNil(t, err)

	mockedUR.On("GetByToken", "token2").Return(user, nil)
	userGot, err = US.GetByToken("token2")
	assert.NotNil(t, userGot)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userGot.ID)
}

func TestUserService_Login_Failed(t *testing.T) {
	mockedUR := &mocks.UserRepository{}
	mockedES := &mocks.EmailSender{}
	mockedUC := &mocks.UserCache{}
	US := NewUserService(mockedUR, mockedES, mockedUC)

	// email does not exist scenario
	mockedUR.On("GetByEmail", "iso@iso.com").Return(nil, errors.New("dummy")).Once()
	token, err := US.Login("iso@iso.com", "123456")
	assert.Nil(t, token)
	assert.Equal(t, domain.ErrorUserNotFound, err)

	// wrong password scenario
	mockedUR.On("GetByEmail", "iso@iso.com").Return(&domain.User{Password: "123"}, nil).Once()
	token, err = US.Login("iso@iso.com", "123456")
	assert.Nil(t, token)
	assert.Equal(t, domain.ErrorWrongPassword, err)
}

func TestUserService_Login_Succeed(t *testing.T) {
	mockedUR := &mocks.UserRepository{}
	mockedES := &mocks.EmailSender{}
	mockedUC := &mocks.UserCache{}
	US := NewUserService(mockedUR, mockedES, mockedUC)

	mockedUR.On("GetByEmail", "iso@iso.com").Return(
		&domain.User{
			Password: generatePassword("123456"),
		},
		nil,
	)
	token, err := US.Login("iso@iso.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, token)

	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%s%s", "iso@iso.com", "123456")))
	expectedToken := hex.EncodeToString(h.Sum(nil))
	assert.Equal(t, expectedToken, *token)
}

func TestUserService_Register_Failed(t *testing.T) {
	mockedUR := &mocks.UserRepository{}
	mockedES := &mocks.EmailSender{}
	mockedUC := &mocks.UserCache{}
	US := NewUserService(mockedUR, mockedES, mockedUC)

	mockedUR.On("GetByEmail", "iso@iso.com").Return(&domain.User{}, nil)
	err := US.Register("iso@iso.com", "123456", "ismail", "bayram")
	assert.True(t, errors.Is(err, domain.ErrorUserAlreadyExists))
	mockedES.AssertNotCalled(t, "SendWelcomeEmail", "iso@iso.com")
}

func TestUserService_Register_Succeed(t *testing.T) {
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
			Password:  generatePassword("123456"),
		},
	).Return(
		&domain.User{
			ID:        3,
			Email:     "iso@iso.com",
			FirstName: "ismail",
			LastName:  "bayram",
			IsActive:  true,
			Password:  generatePassword("123456"),
		},
		nil,
	)
	mockedES.On("SendWelcomeEmail", "iso@iso.com")
	err := US.Register("iso@iso.com", "123456", "ismail", "bayram")
	assert.Nil(t, err)
	time.Sleep(1 * time.Millisecond)
	mockedES.AssertCalled(t, "SendWelcomeEmail", "iso@iso.com")
}

func TestUserService_Verify_Failed(t *testing.T) {
	mockedUR := &mocks.UserRepository{}
	mockedES := &mocks.EmailSender{}
	mockedUC := &mocks.UserCache{}
	US := NewUserService(mockedUR, mockedES, mockedUC)

	mockedUC.On("GetUserIDByVerificationToken", "token").Return(nil).Once()
	err := US.Verify("token")
	assert.Equal(t, domain.ErrorUserNotFound, err)

	cachedID := uint(1)
	mockedUC.On("GetUserIDByVerificationToken", "token").Return(&cachedID)
	mockedUR.On("GetByID", uint(1)).Return(nil, errors.New("dummy"))
	err = US.Verify("token")
	assert.Equal(t, domain.ErrorGeneral, err)
	mockedUR.AssertNotCalled(t, "Update")
}

func TestUserService_Verify_Succeed(t *testing.T) {
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

package services

import (
	"errors"
	"github.com/ismailbayram/shopping/internal/users/models"
	"github.com/ismailbayram/shopping/test/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUserService_GetByID(t *testing.T) {
	user := models.User{
		ID:       2,
		IsActive: true,
	}

	mockedUR := mocks.NewUserRepository(t)
	mockedES := mocks.NewEmailSender(t)
	mockedUC := mocks.NewUserCache(t)
	US := NewUserService(mockedUR, mockedES, mockedUC)

	mockedUR.On("GetByID", uint(1)).Return(models.User{}, models.ErrorUserNotFound)
	userGot, err := US.GetByID(1)
	assert.Equal(t, uint(0), userGot.ID)
	assert.NotNil(t, err)

	mockedUR.On("GetByID", uint(2)).Return(user, nil)
	userGot, err = US.GetByID(2)
	assert.NotNil(t, userGot)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userGot.ID)
}

func TestUserService_GetByToken(t *testing.T) {
	user := models.User{
		ID:       2,
		IsActive: true,
	}

	mockedUR := mocks.NewUserRepository(t)
	mockedES := mocks.NewEmailSender(t)
	mockedUC := mocks.NewUserCache(t)
	US := NewUserService(mockedUR, mockedES, mockedUC)

	mockedUR.On("GetByToken", "token1").Return(models.User{}, models.ErrorUserNotFound)
	userGot, err := US.GetByToken("token1")
	assert.Equal(t, uint(0), userGot.ID)
	assert.NotNil(t, err)

	mockedUR.On("GetByToken", "token2").Return(user, nil)
	userGot, err = US.GetByToken("token2")
	assert.NotEqual(t, uint(0), userGot.ID)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userGot.ID)
}

func TestUserService_Login_Failed(t *testing.T) {
	mockedUR := mocks.NewUserRepository(t)
	mockedES := mocks.NewEmailSender(t)
	mockedUC := mocks.NewUserCache(t)
	US := NewUserService(mockedUR, mockedES, mockedUC)

	// email does not exist scenario
	mockedUR.On("GetByEmail", "iso@iso.com").Return(models.User{}, errors.New("dummy")).Once()
	token, err := US.Login("iso@iso.com", "123456")
	assert.Empty(t, token)
	assert.Equal(t, models.ErrorUserNotFound, err)

	// user not verified scenario
	mockedUR.On("GetByEmail", "iso@iso.com").Return(models.User{IsVerified: false}, nil).Once()
	token, err = US.Login("iso@iso.com", "123456")
	assert.Empty(t, token)
	assert.Equal(t, models.ErrorUserNotVerified, err)

	// wrong password scenario
	mockedUR.On("GetByEmail", "iso@iso.com").Return(models.User{IsVerified: true, Password: "123"}, nil).Once()
	token, err = US.Login("iso@iso.com", "123456")
	assert.Empty(t, token)
	assert.Equal(t, models.ErrorWrongPassword, err)
}

func TestUserService_Login_Succeed(t *testing.T) {
	mockedUR := mocks.NewUserRepository(t)
	mockedES := mocks.NewEmailSender(t)
	mockedUC := mocks.NewUserCache(t)
	US := NewUserService(mockedUR, mockedES, mockedUC)

	user := models.User{
		IsVerified: true,
	}
	user.SetPassword("123456")
	mockedUR.On("GetByEmail", "iso@iso.com").Return(user, nil)
	token, err := US.Login("iso@iso.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestUserService_Register_Failed(t *testing.T) {
	mockedUR := mocks.NewUserRepository(t)
	mockedES := mocks.NewEmailSender(t)
	mockedUC := mocks.NewUserCache(t)
	US := NewUserService(mockedUR, mockedES, mockedUC)

	mockedUR.On("GetByEmail", "iso@iso.com").Return(models.User{ID: 1}, nil).Once()
	err := US.Register("iso@iso.com", "123456", "ismail", "bayram")
	assert.True(t, errors.Is(err, models.ErrorUserAlreadyExists))
	mockedES.AssertNotCalled(t, "SendWelcomeEmail", "iso@iso.com")

	mockedUR.On("GetByEmail", "iso@iso.com").Return(models.User{}, nil).Once()
	user := models.User{
		Email:     "iso@iso.com",
		FirstName: "ismail",
		LastName:  "bayram",
		IsActive:  true,
		Token:     generateToken("iso@iso.com"),
	}
	user.SetPassword("123456")
	mockedUR.On("Create", user).Return(models.User{}, models.ErrorGeneral)
	err = US.Register("iso@iso.com", "123456", "ismail", "bayram")
	assert.NotNil(t, err)
	time.Sleep(1 * time.Millisecond)
	mockedUC.AssertNotCalled(t, "SetUserVerificationToken")
	mockedES.AssertNotCalled(t, "SendWelcomeEmail", "iso@iso.com")
}

func TestUserService_Register_Succeed(t *testing.T) {
	mockedUR := mocks.NewUserRepository(t)
	mockedES := mocks.NewEmailSender(t)
	mockedUC := mocks.NewUserCache(t)
	US := NewUserService(mockedUR, mockedES, mockedUC)

	mockedUR.On("GetByEmail", "iso@iso.com").Return(models.User{}, nil)
	user := models.User{
		Email:     "iso@iso.com",
		FirstName: "ismail",
		LastName:  "bayram",
		IsActive:  true,
		Token:     generateToken("iso@iso.com"),
	}
	user.SetPassword("123456")
	userReturned := user
	userReturned.ID = 3
	mockedUR.On("Create", user).Return(userReturned, nil)
	expectedToken := generateToken("iso@iso.com")
	mockedES.On("SendWelcomeEmail", expectedToken, "iso@iso.com")
	err := US.Register("iso@iso.com", "123456", "ismail", "bayram")
	assert.Nil(t, err)
	time.Sleep(1 * time.Millisecond)
	mockedES.AssertCalled(t, "SendWelcomeEmail", expectedToken, "iso@iso.com")
}

func TestUserService_Verify_Failed(t *testing.T) {
	mockedUR := mocks.NewUserRepository(t)
	mockedES := mocks.NewEmailSender(t)
	mockedUC := mocks.NewUserCache(t)
	US := NewUserService(mockedUR, mockedES, mockedUC)

	mockedUR.On("GetByToken", "token").Return(models.User{}, errors.New("dummy")).Once()
	err := US.Verify("token")
	assert.Equal(t, models.ErrorUserNotFound, err)
	mockedUR.AssertNotCalled(t, "Update")

	mockedUR.On("GetByToken", "token").Return(models.User{ID: 1}, nil).Once()
	mockedUR.On("Update", models.User{ID: 1, IsVerified: true}).Return(models.ErrorGeneral)
	err = US.Verify("token")
	assert.Equal(t, models.ErrorGeneral, err)
}

func TestUserService_Verify_Succeed(t *testing.T) {
	mockedUR := mocks.NewUserRepository(t)
	mockedES := mocks.NewEmailSender(t)
	mockedUC := mocks.NewUserCache(t)
	US := NewUserService(mockedUR, mockedES, mockedUC)

	cachedID := uint(1)
	mockedUR.On("GetByToken", "token").Return(models.User{ID: cachedID}, nil)
	mockedUR.On("Update", models.User{ID: cachedID, IsVerified: true}).Return(nil)
	err := US.Verify("token")
	assert.Nil(t, err)
	mockedUR.AssertCalled(t, "Update", models.User{ID: cachedID, IsVerified: true})
}

func TestUserService_ChangePassword_Failed(t *testing.T) {
	mockedUR := mocks.NewUserRepository(t)
	mockedES := mocks.NewEmailSender(t)
	mockedUC := mocks.NewUserCache(t)
	US := NewUserService(mockedUR, mockedES, mockedUC)

	user := models.User{}
	user.SetPassword("newpassword")
	mockedUR.On("Update", user).Return(models.ErrorGeneral).Once()
	err := US.ChangePassword(user, "newpassword")
	assert.Equal(t, models.ErrorGeneral, err)
}

func TestUserService_ChangePassword_Success(t *testing.T) {
	mockedUR := mocks.NewUserRepository(t)
	mockedES := mocks.NewEmailSender(t)
	mockedUC := mocks.NewUserCache(t)
	US := NewUserService(mockedUR, mockedES, mockedUC)

	user := models.User{}
	user.SetPassword("newpassword")
	mockedUR.On("Update", user).Return(nil).Once()
	err := US.ChangePassword(user, "newpassword")
	assert.Nil(t, err)
}

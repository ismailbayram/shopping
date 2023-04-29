package services

import (
	"errors"
	"github.com/ismailbayram/shopping/internal/users/domain"
	"github.com/ismailbayram/shopping/test/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUserService_GetByID(t *testing.T) {
	user := domain.User{
		ID:       2,
		IsActive: true,
	}

	mockedUR := mocks.NewUserRepository(t)
	mockedES := mocks.NewEmailSender(t)
	mockedUC := mocks.NewUserCache(t)
	US := NewUserService(mockedUR, mockedES, mockedUC)

	mockedUR.On("GetByID", uint(1)).Return(domain.User{}, domain.ErrorUserNotFound)
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
	user := domain.User{
		ID:       2,
		IsActive: true,
	}

	mockedUR := mocks.NewUserRepository(t)
	mockedES := mocks.NewEmailSender(t)
	mockedUC := mocks.NewUserCache(t)
	US := NewUserService(mockedUR, mockedES, mockedUC)

	mockedUR.On("GetByToken", "token1").Return(domain.User{}, domain.ErrorUserNotFound)
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
	mockedUR.On("GetByEmail", "iso@iso.com").Return(domain.User{}, errors.New("dummy")).Once()
	token, err := US.Login("iso@iso.com", "123456")
	assert.Empty(t, token)
	assert.Equal(t, domain.ErrorUserNotFound, err)

	// user not verified scenario
	mockedUR.On("GetByEmail", "iso@iso.com").Return(domain.User{IsVerified: false}, nil).Once()
	token, err = US.Login("iso@iso.com", "123456")
	assert.Empty(t, token)
	assert.Equal(t, domain.ErrorUserNotVerified, err)

	// wrong password scenario
	mockedUR.On("GetByEmail", "iso@iso.com").Return(domain.User{IsVerified: true, Password: "123"}, nil).Once()
	token, err = US.Login("iso@iso.com", "123456")
	assert.Empty(t, token)
	assert.Equal(t, domain.ErrorWrongPassword, err)
}

func TestUserService_Login_Succeed(t *testing.T) {
	mockedUR := mocks.NewUserRepository(t)
	mockedES := mocks.NewEmailSender(t)
	mockedUC := mocks.NewUserCache(t)
	US := NewUserService(mockedUR, mockedES, mockedUC)

	user := domain.User{
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

	mockedUR.On("GetByEmail", "iso@iso.com").Return(domain.User{ID: 1}, nil).Once()
	err := US.Register("iso@iso.com", "123456", "ismail", "bayram")
	assert.True(t, errors.Is(err, domain.ErrorUserAlreadyExists))
	mockedES.AssertNotCalled(t, "SendWelcomeEmail", "iso@iso.com")

	mockedUR.On("GetByEmail", "iso@iso.com").Return(domain.User{}, nil).Once()
	user := domain.User{
		Email:     "iso@iso.com",
		FirstName: "ismail",
		LastName:  "bayram",
		IsActive:  true,
		Token:     generateToken("iso@iso.com"),
	}
	user.SetPassword("123456")
	mockedUR.On("Create", user).Return(domain.User{}, domain.ErrorGeneral)
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

	mockedUR.On("GetByEmail", "iso@iso.com").Return(domain.User{}, nil)
	user := domain.User{
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

	mockedUR.On("GetByToken", "token").Return(domain.User{}, errors.New("dummy")).Once()
	err := US.Verify("token")
	assert.Equal(t, domain.ErrorUserNotFound, err)
	mockedUR.AssertNotCalled(t, "Update")

	mockedUR.On("GetByToken", "token").Return(domain.User{ID: 1}, nil).Once()
	mockedUR.On("Update", domain.User{ID: 1, IsVerified: true}).Return(domain.ErrorGeneral)
	err = US.Verify("token")
	assert.Equal(t, domain.ErrorGeneral, err)
}

func TestUserService_Verify_Succeed(t *testing.T) {
	mockedUR := mocks.NewUserRepository(t)
	mockedES := mocks.NewEmailSender(t)
	mockedUC := mocks.NewUserCache(t)
	US := NewUserService(mockedUR, mockedES, mockedUC)

	cachedID := uint(1)
	mockedUR.On("GetByToken", "token").Return(domain.User{ID: cachedID}, nil)
	mockedUR.On("Update", domain.User{ID: cachedID, IsVerified: true}).Return(nil)
	err := US.Verify("token")
	assert.Nil(t, err)
	mockedUR.AssertCalled(t, "Update", domain.User{ID: cachedID, IsVerified: true})
}

func TestUserService_ChangePassword_Failed(t *testing.T) {
	mockedUR := mocks.NewUserRepository(t)
	mockedES := mocks.NewEmailSender(t)
	mockedUC := mocks.NewUserCache(t)
	US := NewUserService(mockedUR, mockedES, mockedUC)

	user := domain.User{}
	user.SetPassword("newpassword")
	mockedUR.On("Update", user).Return(domain.ErrorGeneral).Once()
	err := US.ChangePassword(user, "newpassword")
	assert.Equal(t, domain.ErrorGeneral, err)
}

func TestUserService_ChangePassword_Success(t *testing.T) {
	mockedUR := mocks.NewUserRepository(t)
	mockedES := mocks.NewEmailSender(t)
	mockedUC := mocks.NewUserCache(t)
	US := NewUserService(mockedUR, mockedES, mockedUC)

	user := domain.User{}
	user.SetPassword("newpassword")
	mockedUR.On("Update", user).Return(nil).Once()
	err := US.ChangePassword(user, "newpassword")
	assert.Nil(t, err)
}

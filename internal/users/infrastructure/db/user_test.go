package infrastructure

import (
	"github.com/ismailbayram/shopping/internal/users/models"
	"github.com/ismailbayram/shopping/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserDBTestSuite struct {
	test.AppTestSuite
}

func (s *UserDBTestSuite) TestCreate() {
	udbr := NewUserDBRepository(s.DB)
	user, err := udbr.Create(models.User{
		Email:      "iso@iso.com",
		FirstName:  "ismail",
		LastName:   "bayram",
		IsActive:   true,
		IsAdmin:    false,
		IsVerified: false,
		Password:   "asdasd",
	})
	assert.Nil(s.T(), err)
	assert.NotEqual(s.T(), uint(0), user.ID)

	user, err = udbr.Create(models.User{
		Email:      "iso@iso.com",
		FirstName:  "ismail",
		LastName:   "bayram",
		IsActive:   true,
		IsAdmin:    false,
		IsVerified: false,
		Password:   "asdasd",
	})
	assert.Equal(s.T(), models.ErrorGeneral, err)
}

func (s *UserDBTestSuite) TestUpdate() {
	udbr := NewUserDBRepository(s.DB)

	created, err := udbr.Create(models.User{Email: "iso@iso.com"})
	assert.Nil(s.T(), err)

	created.Email = "new@iso.com"
	err = udbr.Update(created)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), "new@iso.com", created.Email)

	user, err := udbr.GetByEmail(created.Email)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), created.ID, user.ID)
}

func (s *UserDBTestSuite) TestGetByID() {
	udbr := NewUserDBRepository(s.DB)

	created, err := udbr.Create(models.User{Email: "test@png.com"})
	assert.Nil(s.T(), err)

	user, err := udbr.GetByID(created.ID)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), models.User{ID: user.ID, Email: "test@png.com"}, user)

	notExist, err := udbr.GetByID(0)
	assert.Equal(s.T(), models.ErrorUserNotFound, err)
	assert.Equal(s.T(), models.User{}, notExist)
}

func (s *UserDBTestSuite) TestGetByEmail() {
	udbr := NewUserDBRepository(s.DB)

	created, err := udbr.Create(models.User{Email: "test@png.com"})
	assert.Nil(s.T(), err)

	user, err := udbr.GetByEmail(created.Email)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), models.User{ID: user.ID, Email: "test@png.com"}, user)

	notExist, err := udbr.GetByEmail("none@mail.com")
	assert.Equal(s.T(), models.ErrorUserNotFound, err)
	assert.Equal(s.T(), models.User{}, notExist)
}

func (s *UserDBTestSuite) TestGetByToken() {
	udbr := NewUserDBRepository(s.DB)

	created, err := udbr.Create(models.User{Token: "token"})
	assert.Nil(s.T(), err)

	user, err := udbr.GetByToken(created.Token)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), models.User{ID: user.ID, Token: "token"}, user)

	notExist, err := udbr.GetByToken("none@mail.com")
	assert.Equal(s.T(), models.ErrorUserNotFound, err)
	assert.Equal(s.T(), models.User{}, notExist)
}

func (s *UserDBTestSuite) TestAll() {
	udbr := NewUserDBRepository(s.DB)

	created1, err := udbr.Create(models.User{Email: "test1@png.com", Token: "1"})
	assert.Nil(s.T(), err)
	_, err = udbr.Create(models.User{Email: "test2@png.com", Token: "2"})
	assert.Nil(s.T(), err)

	users, err := udbr.All(map[string]interface{}{})
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 2, len(users))
	assert.Equal(s.T(), created1.ID, users[0].ID)

	users, err = udbr.All(map[string]interface{}{"email": "test1@png.com"})
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 1, len(users))

	users, err = udbr.All(map[string]interface{}{"email": "none@none.com"})
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 0, len(users))
}

func TestUserDBTestSuite(t *testing.T) {
	userDbTestSuite := new(UserDBTestSuite)
	userDbTestSuite.Models = []interface{}{&UserDB{}}
	suite.Run(t, userDbTestSuite)
}

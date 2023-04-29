package infrastructure

import (
	"github.com/ismailbayram/shopping/internal/users/models"
	"time"
)

func ToUser(userDB UserDB) models.User {
	return models.User{
		ID:         userDB.ID,
		Email:      userDB.Email,
		FirstName:  userDB.FirstName,
		LastName:   userDB.LastName,
		IsActive:   userDB.IsActive,
		IsVerified: userDB.IsVerified,
		IsAdmin:    userDB.IsAdmin,
		Password:   userDB.Password,
		Token:      userDB.Token,
	}
}

func ToUserDB(user models.User) UserDB {
	return UserDB{
		ID:         user.ID,
		UpdatedAt:  time.Now(),
		Email:      user.Email,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		IsActive:   user.IsActive,
		IsVerified: user.IsVerified,
		IsAdmin:    user.IsAdmin,
		Password:   user.Password,
		Token:      user.Token,
	}
}

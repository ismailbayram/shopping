package infrastructure

import (
	"github.com/ismailbayram/shopping/internal/users/domain"
	"time"
)

func ToUser(userDB UserDB) domain.User {
	return domain.User{
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

func ToUserDB(user domain.User) UserDB {
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

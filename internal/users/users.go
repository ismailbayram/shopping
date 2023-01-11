package users

import (
	"github.com/ismailbayram/shopping/internal/users/application"
	infrastructure "github.com/ismailbayram/shopping/internal/users/infrastructure/db"
	"gorm.io/gorm"
)

type Users struct {
	Service *application.UserService
}

func New(db *gorm.DB) Users {
	return Users{
		Service: application.NewUserService(
			&infrastructure.UserDBRepository{DB: db},
			nil,
			nil,
		),
	}
}

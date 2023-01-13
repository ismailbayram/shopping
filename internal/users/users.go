package users

import (
	infrastructure "github.com/ismailbayram/shopping/internal/users/infrastructure/db"
	"github.com/ismailbayram/shopping/internal/users/services"
	"gorm.io/gorm"
)

type Users struct {
	Service *services.UserService
}

func New(db *gorm.DB) Users {
	return Users{
		Service: services.NewUserService(
			&infrastructure.UserDBRepository{DB: db},
			nil,
			nil,
		),
	}
}

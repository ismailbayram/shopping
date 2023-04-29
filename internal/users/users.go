package users

import (
	"github.com/ismailbayram/shopping/config"
	infrastructureCache "github.com/ismailbayram/shopping/internal/users/infrastructure/cache"
	infrastructureDB "github.com/ismailbayram/shopping/internal/users/infrastructure/db"
	infrastructureSMTP "github.com/ismailbayram/shopping/internal/users/infrastructure/smtp"
	"github.com/ismailbayram/shopping/internal/users/presentation"
	"github.com/ismailbayram/shopping/internal/users/services"
	"gorm.io/gorm"
)

type Users struct {
	Views   presentation.UserViews
	Service *services.UserService
}

func New(db *gorm.DB, cfg *config.Configuration) Users {
	userCache := infrastructureCache.NewUserCache()
	emailSender := infrastructureSMTP.NewEmailSender(cfg.SMTP)

	userService := services.NewUserService(
		infrastructureDB.NewUserDBRepository(db),
		emailSender,
		userCache,
	)

	return Users{
		Views:   presentation.NewUserViews(userService),
		Service: userService,
	}
}

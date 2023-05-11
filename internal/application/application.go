package application

import (
	"github.com/ismailbayram/shopping/internal/media"
	"github.com/ismailbayram/shopping/internal/users"
)

type Application struct {
	SiteUrl string
	Users   users.Users
	Media   media.Media
}

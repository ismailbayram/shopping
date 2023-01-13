package application

import (
	"github.com/ismailbayram/shopping/internal/media"
	"github.com/ismailbayram/shopping/internal/users"
)

type Application struct {
	Users users.Users
	Media media.Media
}

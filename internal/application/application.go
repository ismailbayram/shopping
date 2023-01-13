package application

import (
	"fmt"
	"github.com/ismailbayram/shopping/internal/media"
	"github.com/ismailbayram/shopping/internal/users"
)

type Application struct {
	SiteUrl  string
	MediaUrl string
	Users    users.Users
	Media    media.Media
}

func (app *Application) GetMediaUrl(filePath string) string {
	return fmt.Sprintf("%s%s/%s", app.SiteUrl, app.MediaUrl, filePath)
}

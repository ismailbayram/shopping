package media

import (
	"github.com/ismailbayram/shopping/internal/media/application"
	infrastructure "github.com/ismailbayram/shopping/internal/media/infrastructure/storage"
	"gorm.io/gorm"
)

type Media struct {
	Service *application.ImageService
}

func New(db *gorm.DB, mediaRoot string) Media {
	return Media{
		Service: application.NewImageService(
			nil,
			infrastructure.NewFileStorage(mediaRoot),
		),
	}
}

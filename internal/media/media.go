package media

import (
	dbInfrastructure "github.com/ismailbayram/shopping/internal/media/infrastructure/db"
	fileInfrastructure "github.com/ismailbayram/shopping/internal/media/infrastructure/storage"
	"github.com/ismailbayram/shopping/internal/media/services"
	"gorm.io/gorm"
)

type Media struct {
	Service services.ImageService
}

func New(db *gorm.DB, mediaRoot string) Media {
	return Media{
		Service: services.NewImageService(
			dbInfrastructure.NewImageDBRepository(db),
			fileInfrastructure.NewFileStorage(mediaRoot),
		),
	}
}

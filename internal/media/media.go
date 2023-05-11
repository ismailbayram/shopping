package media

import (
	dbInfrastructure "github.com/ismailbayram/shopping/internal/media/infrastructure/db"
	fileInfrastructure "github.com/ismailbayram/shopping/internal/media/infrastructure/storage"
	"github.com/ismailbayram/shopping/internal/media/presentation"
	"github.com/ismailbayram/shopping/internal/media/services"
	"gorm.io/gorm"
)

type Media struct {
	Views        presentation.MediaViews
	ImageService services.ImageService
}

func New(db *gorm.DB, mediaRoot string, mediaUrl string) Media {
	imageService := services.NewImageService(
		dbInfrastructure.NewImageDBRepository(db),
		fileInfrastructure.NewFileStorage(mediaRoot, mediaUrl),
	)

	return Media{
		Views:        presentation.NewMediaViews(imageService),
		ImageService: imageService,
	}
}

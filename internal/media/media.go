package media

import (
	dbInfrastructure "github.com/ismailbayram/shopping/internal/media/infrastructure/db"
	fileInfrastructure "github.com/ismailbayram/shopping/internal/media/infrastructure/storage"
	"github.com/ismailbayram/shopping/internal/media/presentation"
	"github.com/ismailbayram/shopping/internal/media/services"
	"gorm.io/gorm"
)

type Media struct {
	Views presentation.MediaViews
}

func New(db *gorm.DB, mediaRoot string) Media {
	return Media{
		Views: presentation.NewMediaViews(
			services.NewImageService(
				dbInfrastructure.NewImageDBRepository(db),
				fileInfrastructure.NewFileStorage(mediaRoot),
			),
		),
	}
}

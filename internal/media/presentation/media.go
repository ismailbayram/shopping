package presentation

import (
	"github.com/ismailbayram/shopping/internal/media/models"
)

type ImageService interface {
	GetByID(uint) (models.Image, error)
	Create(string, []byte) (models.Image, error)
}

type MediaViews struct {
	Service ImageService
}

func NewMediaViews(service ImageService) MediaViews {
	return MediaViews{
		Service: service,
	}
}

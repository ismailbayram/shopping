package presentation

import (
	"fmt"
	"github.com/ismailbayram/shopping/internal/media/models"
	"github.com/spf13/viper"
)

type ImageService interface {
	GetByID(uint) (models.Image, error)
	Create(string, []byte) (models.Image, error)
}

type MediaViews struct {
	Service ImageService
}

func (view *MediaViews) GetBaseURL() string {
	return fmt.Sprintf("%s/%s", viper.GetString("server.domain"), viper.GetString("server.mediaurl"))
}

func NewMediaViews(service ImageService) MediaViews {
	return MediaViews{
		Service: service,
	}
}

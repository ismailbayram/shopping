package presentation

import (
	"fmt"
	domain "github.com/ismailbayram/shopping/internal/media/domain/models"
	"github.com/spf13/viper"
)

type ImageService interface {
	GetByID(uint) (domain.Image, error)
	Create(string, []byte) (domain.Image, error)
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

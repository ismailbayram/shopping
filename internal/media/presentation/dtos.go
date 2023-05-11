package presentation

import (
	"github.com/ismailbayram/shopping/internal/media/models"
)

type ImageDTO struct {
	ID  int    `json:"id"`
	Url string `json:"url"`
}

func ToImageDTO(image models.Image) ImageDTO {
	return ImageDTO{
		ID:  int(image.ID),
		Url: image.URL,
	}
}

package presentation

import domain "github.com/ismailbayram/shopping/internal/media/domain/models"

type ImageDTO struct {
	ID  int    `json:"id"`
	Url string `json:"url"`
}

func ToImageDTO(baseUrl string, image domain.Image) ImageDTO {
	return ImageDTO{
		ID:  int(image.ID),
		Url: image.Url(baseUrl),
	}
}

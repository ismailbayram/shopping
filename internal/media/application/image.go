package application

import domain "github.com/ismailbayram/shopping/internal/media/domain/models"

type ImageRepository interface {
	Create(domain.Image) (domain.Image, error)
	GetByID(uint) (domain.Image, error)
}

type ImageService struct {
	repo ImageRepository
}

func NewImageService(repo ImageRepository) *ImageService {
	return &ImageService{repo: repo}
}

func (is *ImageService) GetByID(id uint) (domain.Image, error) {
	return is.repo.GetByID(id)
}

// TODO: check same file names.

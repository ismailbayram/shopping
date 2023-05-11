package services

import (
	"github.com/ismailbayram/shopping/internal/media/models"
)

type ImageRepository interface {
	Create(models.Image) (models.Image, error)
	GetByID(uint) (models.Image, error)
}

type ImageStorage interface {
	Upload(string, []byte) (string, error)
	Url(string) string
}

type ImageService struct {
	repo    ImageRepository
	storage ImageStorage
}

func NewImageService(repo ImageRepository, storage ImageStorage) ImageService {
	return ImageService{
		repo:    repo,
		storage: storage,
	}
}

func (is ImageService) GetByID(id uint) (models.Image, error) {
	image, err := is.repo.GetByID(id)

	if err != nil {
		return models.Image{}, err
	}
	image.URL = is.storage.Url(image.Path)
	return image, nil
}

func (is ImageService) Create(name string, content []byte) (models.Image, error) {
	path, err := is.storage.Upload(name, content)
	if err != nil {
		return models.Image{}, models.ErrorGeneral
	}

	image, err := is.repo.Create(models.Image{Path: path})
	if err != nil {
		return models.Image{}, models.ErrorGeneral
	}

	image.URL = is.storage.Url(image.Path)

	return image, nil
}

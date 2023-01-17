package services

import domain "github.com/ismailbayram/shopping/internal/media/domain/models"

type ImageRepository interface {
	Create(domain.Image) (domain.Image, error)
	GetByID(uint) (domain.Image, error)
}

type ImageStorage interface {
	Upload(string, []byte) (string, error)
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

func (is ImageService) GetByID(id uint) (domain.Image, error) {
	return is.repo.GetByID(id)
}

func (is ImageService) Create(name string, content []byte) (domain.Image, error) {
	path, err := is.storage.Upload(name, content)
	if err != nil {
		return domain.Image{}, domain.ErrorGeneral
	}

	return is.repo.Create(domain.Image{Path: path})
}

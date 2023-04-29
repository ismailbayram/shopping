package services

import (
	"github.com/ismailbayram/shopping/internal/products/models"
)

type CategoryRepository interface {
	GetByID(uint) (models.Category, error)
	Create(models.Category) (models.Category, error)
	Update(models.Category) error
	All() ([]models.Category, error)
}

type CategoryService struct {
	repo        CategoryRepository
	productRepo ProductRepository
}

func NewCategoryService(repo CategoryRepository, productRepo ProductRepository) *CategoryService {
	return &CategoryService{repo: repo, productRepo: productRepo}
}

func (cs *CategoryService) GetByID(id uint) (models.Category, error) {
	return cs.repo.GetByID(id)
}

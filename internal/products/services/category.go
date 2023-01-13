package services

import domain "github.com/ismailbayram/shopping/internal/products/domain/models"

type CategoryRepository interface {
	GetByID(uint) (domain.Category, error)
	Create(domain.Category) (domain.Category, error)
	Update(domain.Category) error
	All() ([]domain.Category, error)
}

type CategoryService struct {
	repo        CategoryRepository
	productRepo ProductRepository
}

func NewCategoryService(repo CategoryRepository, productRepo ProductRepository) *CategoryService {
	return &CategoryService{repo: repo, productRepo: productRepo}
}

func (cs *CategoryService) GetByID(id uint) (domain.Category, error) {
	return cs.repo.GetByID(id)
}

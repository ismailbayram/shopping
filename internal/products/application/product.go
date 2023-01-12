package application

import domain "github.com/ismailbayram/shopping/internal/products/domain/models"

type ProductRepository interface {
	GetByID(uint) (domain.Product, error)
	Create(domain.Product) (domain.Product, error)
	Update(domain.Product) error
	All() ([]domain.Product, error)
	GetByCategory(category domain.Category) ([]domain.Product, error)
}

type ProductService struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (ps *ProductService) GetByID(id uint) (domain.Product, error) {
	return ps.repo.GetByID(id)
}

package services

import (
	domain2 "github.com/ismailbayram/shopping/internal/products/models"
)

type ProductRepository interface {
	GetByID(uint) (domain2.Product, error)
	Create(domain2.Product) (domain2.Product, error)
	Update(domain2.Product) error
	All() ([]domain2.Product, error)
	GetByCategory(category domain2.Category) ([]domain2.Product, error)
}

type ProductService struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (ps *ProductService) GetByID(id uint) (domain2.Product, error) {
	return ps.repo.GetByID(id)
}

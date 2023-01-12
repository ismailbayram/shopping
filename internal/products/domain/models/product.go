package domain

import "errors"

var (
	ErrorGeneral         = errors.New("Something went wrong, please try again.")
	ErrorProductNotFound = errors.New("Product not found.")
)

type Product struct {
	ID       uint
	Name     string
	Stock    uint
	Price    float32
	Category Category
}

func (p *Product) IsAvailable() bool {
	return p.Stock > 0
}

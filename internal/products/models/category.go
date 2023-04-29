package models

import "errors"

var (
	ErrorCategoryNotFound = errors.New("Category not found.")
)

type Category struct {
	ID       uint
	Name     string
	Products []Product
}

func (c *Category) IsEmpty() bool {
	return len(c.Products) == 0
}

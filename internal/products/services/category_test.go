package services

import (
	domain "github.com/ismailbayram/shopping/internal/products/domain/models"
	"github.com/ismailbayram/shopping/test/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategoryService_GetByID(t *testing.T) {
	category := domain.Category{
		ID:   2,
		Name: "test",
	}

	mockedCR := &mocks.CategoryRepository{}
	mockedPR := &mocks.ProductRepository{}
	PS := NewCategoryService(mockedCR, mockedPR)

	mockedCR.On("GetByID", uint(1)).Return(domain.Category{}, domain.ErrorCategoryNotFound)
	categoryGot, err := PS.GetByID(1)
	assert.Equal(t, uint(0), categoryGot.ID)
	assert.NotNil(t, err)

	mockedCR.On("GetByID", uint(2)).Return(category, nil)
	categoryGot, err = PS.GetByID(2)
	assert.NotEqual(t, uint(0), categoryGot)
	assert.Nil(t, err)
	assert.Equal(t, category.ID, categoryGot.ID)
}

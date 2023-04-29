package services

import (
	"github.com/ismailbayram/shopping/internal/products/models"
	"github.com/ismailbayram/shopping/test/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductService_GetByID(t *testing.T) {
	product := models.Product{
		ID:   2,
		Name: "test",
	}

	mockedPR := &mocks.ProductRepository{}
	PS := NewProductService(mockedPR)

	mockedPR.On("GetByID", uint(1)).Return(models.Product{}, models.ErrorProductNotFound)
	productGot, err := PS.GetByID(1)
	assert.Equal(t, uint(0), productGot.ID)
	assert.NotNil(t, err)

	mockedPR.On("GetByID", uint(2)).Return(product, nil)
	productGot, err = PS.GetByID(2)
	assert.NotEqual(t, uint(0), productGot.ID)
	assert.Nil(t, err)
	assert.Equal(t, product.ID, productGot.ID)
}

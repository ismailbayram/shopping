package application

import (
	domain "github.com/ismailbayram/shopping/internal/products/domain/models"
	"github.com/ismailbayram/shopping/test/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductService_GetByID(t *testing.T) {
	product := &domain.Product{
		ID:   2,
		Name: "test",
	}

	mockedPR := &mocks.ProductRepository{}
	PS := NewProductService(mockedPR)

	mockedPR.On("GetByID", uint(1)).Return(nil, domain.ErrorProductNotFound)
	productGot, err := PS.GetByID(1)
	assert.Nil(t, productGot)
	assert.NotNil(t, err)

	mockedPR.On("GetByID", uint(2)).Return(product, nil)
	productGot, err = PS.GetByID(2)
	assert.NotNil(t, productGot)
	assert.Nil(t, err)
	assert.Equal(t, product.ID, productGot.ID)
}

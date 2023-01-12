// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	models "github.com/ismailbayram/shopping/internal/products/domain/models"
	mock "github.com/stretchr/testify/mock"
)

// ProductRepository is an autogenerated mock type for the ProductRepository type
type ProductRepository struct {
	mock.Mock
}

// All provides a mock function with given fields:
func (_m *ProductRepository) All() ([]models.Product, error) {
	ret := _m.Called()

	var r0 []models.Product
	if rf, ok := ret.Get(0).(func() []models.Product); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: _a0
func (_m *ProductRepository) Create(_a0 *models.Product) (*models.Product, error) {
	ret := _m.Called(_a0)

	var r0 *models.Product
	if rf, ok := ret.Get(0).(func(*models.Product) *models.Product); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Product) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByCategory provides a mock function with given fields: category
func (_m *ProductRepository) GetByCategory(category models.Category) ([]models.Product, error) {
	ret := _m.Called(category)

	var r0 []models.Product
	if rf, ok := ret.Get(0).(func(models.Category) []models.Product); ok {
		r0 = rf(category)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Category) error); ok {
		r1 = rf(category)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: _a0
func (_m *ProductRepository) GetByID(_a0 uint) (*models.Product, error) {
	ret := _m.Called(_a0)

	var r0 *models.Product
	if rf, ok := ret.Get(0).(func(uint) *models.Product); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0
func (_m *ProductRepository) Update(_a0 *models.Product) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Product) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewProductRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewProductRepository creates a new instance of ProductRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProductRepository(t mockConstructorTestingTNewProductRepository) *ProductRepository {
	mock := &ProductRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

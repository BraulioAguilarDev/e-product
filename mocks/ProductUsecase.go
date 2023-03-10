// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	model "ecommerce/model"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// ProductUsecase is an autogenerated mock type for the ProductUsecase type
type ProductUsecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: sku, name, price
func (_m *ProductUsecase) Create(sku string, name string, price float64) (*model.Product, error) {
	ret := _m.Called(sku, name, price)

	var r0 *model.Product
	if rf, ok := ret.Get(0).(func(string, string, float64) *model.Product); ok {
		r0 = rf(sku, name, price)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, float64) error); ok {
		r1 = rf(sku, name, price)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *ProductUsecase) Delete(id uuid.UUID) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: id
func (_m *ProductUsecase) Get(id uuid.UUID) (*model.Product, error) {
	ret := _m.Called(id)

	var r0 *model.Product
	if rf, ok := ret.Get(0).(func(uuid.UUID) *model.Product); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: product
func (_m *ProductUsecase) Update(product *model.Product) (*model.Product, error) {
	ret := _m.Called(product)

	var r0 *model.Product
	if rf, ok := ret.Get(0).(func(*model.Product) *model.Product); ok {
		r0 = rf(product)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Product) error); ok {
		r1 = rf(product)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewProductUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewProductUsecase creates a new instance of ProductUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProductUsecase(t mockConstructorTestingTNewProductUsecase) *ProductUsecase {
	mock := &ProductUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

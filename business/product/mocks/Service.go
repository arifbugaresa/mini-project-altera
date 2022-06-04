// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	request "github.com/arifbugaresa/mini-project-altera/api/v1/product/request"
	product "github.com/arifbugaresa/mini-project-altera/business/product"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// DeleteProductByID provides a mock function with given fields: id
func (_m *Service) DeleteProductByID(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAllProduct provides a mock function with given fields:
func (_m *Service) FindAllProduct() (product.Database, error) {
	ret := _m.Called()

	var r0 product.Database
	if rf, ok := ret.Get(0).(func() product.Database); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(product.Database)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindProductByID provides a mock function with given fields: id
func (_m *Service) FindProductByID(id int) (product.Product, error) {
	ret := _m.Called(id)

	var r0 product.Product
	if rf, ok := ret.Get(0).(func(int) product.Product); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(product.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertProduct provides a mock function with given fields: userParam
func (_m *Service) InsertProduct(userParam request.InsertProductRequest) error {
	ret := _m.Called(userParam)

	var r0 error
	if rf, ok := ret.Get(0).(func(request.InsertProductRequest) error); ok {
		r0 = rf(userParam)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type NewServiceT interface {
	mock.TestingT
	Cleanup(func())
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t NewServiceT) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
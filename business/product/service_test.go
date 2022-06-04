package product_test

import (
	"errors"
	"github.com/arifbugaresa/mini-project-altera/api/v1/product/request"
	"github.com/arifbugaresa/mini-project-altera/business/product"
	"github.com/arifbugaresa/mini-project-altera/business/product/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestInsert(t *testing.T)  {
	mockProductRepo := new(mocks.Repository)
	
	// data product palsu
	mockProductData := request.InsertProductRequest{
		Name:      "Product 1",
		Price:     2000,
		CreatedBy: "Testing",
	}

	// Positive Case tanpa error
	t.Run("success", func(t *testing.T) {
		mockProductRepo.On("InsertProduct", mock.Anything).Return(nil).Once()

		productService := product.NewService(mockProductRepo)

		err := productService.InsertProduct(mockProductData)


		assert.NoError(t, err)
	})

	//Negative Test menghasilkan error
	t.Run("failed", func(t *testing.T) {
		mockProductRepo.On("InsertProduct", mock.Anything).Return(errors.New("error unit test")).Once()

		productService := product.NewService(mockProductRepo)

		err := productService.InsertProduct(mockProductData)

		assert.Error(t, err)
	})

}
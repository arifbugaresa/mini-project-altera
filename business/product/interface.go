package product

import "github.com/arifbugaresa/mini-project-altera/api/v1/product/request"

type Service interface {
	InsertProduct(userParam request.InsertProductRequest) error
}

type Repository interface {
	SaveProduct(product Product) error
}
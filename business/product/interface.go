package product

import "github.com/arifbugaresa/mini-project-altera/api/v1/product/request"

type Service interface {
	InsertProduct(userParam request.InsertProductRequest) error
	FindProductByID(id int) (Product, error)
	FindAllProduct() (Database, error)
	DeleteProductByID(id int) error
}

type Repository interface {
	InsertProduct(product Product) error
	FindProductByID(id int) (Product, error)
	FindAllProduct() ([]Product, error)
	DeleteProductByID(id int) error
}
package product

import (
	"github.com/arifbugaresa/mini-project-altera/business/product"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db,
	}
}

func (r *repository) SaveProduct(product product.Product) error {
	return r.db.Create(&product).Error
}

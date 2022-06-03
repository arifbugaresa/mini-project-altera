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

func (r *repository) InsertProduct(product product.Product) error {
	return r.db.Create(&product).Error
}

func (r *repository) FindProductByID(id int) (product.Product, error) {
	var product product.Product
	err := r.db.First(&product, id).Error
	return product, err
}

func (r *repository) FindAllProduct() ([]product.Product, error) {
	var listProduct []product.Product
	err := r.db.Find(&listProduct).Error
	return listProduct, err
}

func (r *repository) DeleteProductByID(id int) error {
	return r.db.Delete(&product.Product{}, id).Error
}

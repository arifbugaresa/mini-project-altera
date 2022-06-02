package product

import (
	"gorm.io/gorm"
)

// Product model for gorm make table on db
type Product struct {
	gorm.Model
	Name       string
	Price      uint
	CreatedBy  string
	ModifiedBy string
}

// Database is used to store product list
type Database struct {
	Products []Product
}

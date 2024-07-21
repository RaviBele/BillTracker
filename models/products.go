package models

import "github.com/google/uuid"
import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID       uuid.UUID     `gorm:"type:char(36);primaryKey" json:"id"`
	product  VendorProduct `gorm:"foreignkey:VendorProductID" json:"product"`
	quantity int           `gorm:"not null" json:"quantity"`
}

func NewProduct(product VendorProduct, quantity int) *Product {
	return &Product{product: product, quantity: quantity}
}

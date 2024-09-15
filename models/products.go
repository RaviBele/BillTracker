package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID              uuid.UUID     `gorm:"type:char(36);primaryKey" json:"id"`
	Product         VendorProduct `gorm:"foreignkey:VendorProductID" json:"product"`
	VendorProductID uuid.UUID     `gorm:"type:uuid" json:"vendor_product_id"`
	Quantity        int           `gorm:"null" json:"quantity"`
	gorm.Model
}

func NewProduct(product VendorProduct, quantity int) *Product {
	return &Product{Product: product, Quantity: quantity}
}

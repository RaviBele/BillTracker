package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Vendor struct {
	gorm.Model
	ID       uuid.UUID       `gorm:"type:char(36);primaryKey" json:"id"`
	Name     string          `gorm:"column:name;size:255;not null;" json:"name"`
	Products []VendorProduct `gorm:"foreignKey:VendorID" json:"products"`
}

type VendorProduct struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Code     int       `gorm:"unique" json:"code"`
	VendorID uuid.UUID `gorm:"type:uuid;not null" json:"vendor_id"`
	Name     string    `gorm:"column:name;size:255;not null;" json:"name"`
	Rate     float64   `gorm:"type:decimal(10,2)" json:"rate"`
	Tax      float64   `gorm:"type:decimal(5,2);default:0" json:"tax"`
	Vendor   Vendor
}

func NewVendor(name string) *Vendor {
	id, _ := uuid.NewUUID()
	return &Vendor{
		ID:   id,
		Name: name,
	}
}

func NewVendorProduct(name string, code int, vendorID uuid.UUID, rate float64) *VendorProduct {
	id, _ := uuid.NewUUID()
	return &VendorProduct{
		ID:       id,
		Code:     code,
		VendorID: vendorID,
		Name:     name,
		Rate:     rate,
	}
}

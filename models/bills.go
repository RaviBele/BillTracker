package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BillStatus string

type BillType string

const (
	PAID BillStatus = "PAID"
	DUE  BillStatus = "DUE"
)

const (
	ONETIME   BillType = "ONETIME"
	RECURRING BillType = "RECURRING"
)

type Bill struct {
	gorm.Model
	ID        uuid.UUID  `gorm:"type:char(36);primaryKey" json:"id"`
	Amount    float64    `gorm:"column:amount;not null" json:"amount"`
	Products  []Product  `gorm:"many2many:bill_products;" json:"products"`
	VendorID  uuid.UUID  `gorm:"type:uuid" json:"vendor_id"`
	Vendor    Vendor     `gorm:"foreignkey:VendorID"`
	Status    BillStatus `gorm:"column:status;size:255;not null;" json:"status"`
	AccountID uuid.UUID  `gorm:"type:uuid;not null" json:"account_id"`
	BillType  BillType   `gorm:"size:255;not null;" json:"bill_type"`
	TotalTax  float64    `gorm:"column:total_tax;not null" json:"total_tax"`
}

type BillProduct struct {
	BillID    uuid.UUID `gorm:"type:char(36)" json:"bill_id"`
	ProductID uuid.UUID `gorm:"type:char(36)" json:"product_id"`
}

func NewBill(amount float64, products []Product, vendorID uuid.UUID, status BillStatus) *Bill {
	id, _ := uuid.NewUUID()
	return &Bill{
		ID:       id,
		Amount:   amount,
		Products: products,
		VendorID: vendorID,
		Status:   status,
	}
}

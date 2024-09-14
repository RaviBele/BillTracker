package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountType string

const (
	USER   AccountType = "USER"
	VENDOR AccountType = "VENDOR"
)

type Account struct {
	gorm.Model
	ID    uuid.UUID   `gorm:"type:char(36);primaryKey" json:"id"`
	Name  string      `gorm:"column:name;size:255;not null;" json:"name" validate:"required"`
	Email string      `gorm:"size:255;not null;unique" json:"email" validate:"required,email"`
	Phone string      `gorm:"size:20;not null;unique" json:"phone" validate:"required,phone"`
	Bills []Bill      `gorm:"foreignkey:AccountID" json:"bills"`
	Type  AccountType `gorm:"column:acc_type;size:255;not null;" json:"type" validate:"required"`
}

func NewAccount(name string, aType AccountType) *Account {
	id, _ := uuid.NewUUID()
	return &Account{
		ID:   id,
		Name: name,
		Type: aType,
	}
}

type AccountListResponse struct {
	Pagination
	Data []Account `json:"data`
}

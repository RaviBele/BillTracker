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
	Name  string      `gorm:"column:name;size:255;not null;" json:"name"`
	Bills []Bill      `gorm:"foreignkey:AccountID" json:"bills"`
	Type  AccountType `gorm:"column:acc_type;size:255;not null;" json:"type"`
}

func NewAccount(name string, a_type AccountType) *Account {
	id, _ := uuid.NewUUID()
	return &Account{
		ID:   id,
		Name: name,
		Type: a_type,
	}
}

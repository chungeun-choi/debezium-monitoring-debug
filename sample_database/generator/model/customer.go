package model

import "gorm.io/gorm"

// Customer 모델
type Customer struct {
	gorm.Model
	ID          uint64 `gorm:"primary_key;autoIncrement"`
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Address     string
	Orders      []Order `gorm:"foreignKey:CustomerID"`
}

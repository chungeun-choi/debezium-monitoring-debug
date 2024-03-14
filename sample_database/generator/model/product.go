package model

import "gorm.io/gorm"

// Product 모델
type Product struct {
	gorm.Model
	ID          uint64 `gorm:"primary_key;autoIncrement"`
	Name        string
	Description string
	Price       float64
}

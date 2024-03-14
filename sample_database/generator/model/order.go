package model

import (
	"gorm.io/gorm"
	"time"
)

// Order 모델
type Order struct {
	gorm.Model
	ID         uint64 `gorm:"primary_key;autoIncrement"`
	CustomerID uint64
	ProductID  uint64
	OrderDate  time.Time
	Quantity   int
	Product    Product `gorm:"foreignKey:ProductID"`
}

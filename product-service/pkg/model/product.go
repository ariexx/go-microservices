package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name           string
	Banner         string
	ProductDetails []ProductDetail `gorm:"foreignKey:ProductID"`
}

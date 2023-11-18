package model

import "gorm.io/gorm"

type ProductDetail struct {
	gorm.Model
	ProductID uint `gorm:"index"`
	Name      string
	Price     uint32
}

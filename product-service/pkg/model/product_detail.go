package model

import "gorm.io/gorm"

type ProductDetail struct {
	gorm.Model
	ProductID uint    `gorm:"index"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Name      string
	Price     uint32
}

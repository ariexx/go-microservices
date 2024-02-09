package model

import "gorm.io/gorm"

type PaymentChannel struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255);not null"`
	Banner      string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:varchar(255);not null"`
}

func (*PaymentChannel) TableName() string {
	return "payment_channels"
}

package entity

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	PlayerID  string `json:"player_id"`
	Email     string `json:"email"`
	OrderID   string `json:"order_id"`
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Price     int    `json:"price"`
	Total     int    `json:"total"`
}

type CreateOrder struct {
	gorm.Model
	PlayerID  string `json:"player_id"`
	Email     string `json:"email"`
	OrderID   string `json:"order_id"`
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Price     int    `json:"price"`
	Total     int    `json:"total"`
}

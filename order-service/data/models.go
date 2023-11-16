package data

import "gorm.io/gorm"

type orderConfig struct {
	db *gorm.DB
}

// OrderConfig interface
type OrderConfig interface {
	AutoMigrate() error
	Create(order *Order) error
	Find(order *Order) error
	FindByID(order *Order) error
	Delete(order *Order) error
}

func NewOrderConfig(db *gorm.DB) OrderConfig {
	return &orderConfig{db}
}

// Order model
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

// using gorm auto migrate to create tables
func (u *orderConfig) AutoMigrate() error {
	err := u.db.AutoMigrate(&Order{})
	if err != nil {
		return err
	}

	return nil
}

// using gorm create to insert data
func (u *orderConfig) Create(order *Order) error {
	err := u.db.Create(order).Error
	if err != nil {
		return err
	}

	return nil
}

// using gorm find to get data
func (u *orderConfig) Find(order *Order) error {
	err := u.db.Find(order).Error
	if err != nil {
		return err
	}

	return nil
}

// using gorm find to get data
func (u *orderConfig) FindByID(order *Order) error {
	err := u.db.Where("id = ?", order.ID).First(order).Error
	if err != nil {
		return err
	}

	return nil
}

// delete order
func (u *orderConfig) Delete(order *Order) error {
	err := u.db.Delete(order).Error
	if err != nil {
		return err
	}

	return nil
}

package data

import "gorm.io/gorm"

type userConfig struct {
	db *gorm.DB
}

// OrderConfig interface
type OrderConfig interface {
	AutoMigrate() error
	Create(user *Order) error
	Find(user *Order) error
	FindByID(user *Order) error
	Delete(user *Order) error
}

func NewOrderConfig(db *gorm.DB) OrderConfig {
	return &userConfig{db}
}

// Order model
type Order struct {
	gorm.Model
	Email     string `json:"email"`
	OrderID   string `json:"order_id"`
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Price     int    `json:"price"`
	Total     int    `json:"total"`
}

// using gorm auto migrate to create tables
func (u *userConfig) AutoMigrate() error {
	err := u.db.AutoMigrate(&Order{})
	if err != nil {
		return err
	}

	return nil
}

// using gorm create to insert data
func (u *userConfig) Create(user *Order) error {
	err := u.db.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

// using gorm find to get data
func (u *userConfig) Find(user *Order) error {
	err := u.db.Find(user).Error
	if err != nil {
		return err
	}

	return nil
}

// using gorm find to get data
func (u *userConfig) FindByID(user *Order) error {
	err := u.db.Where("id = ?", user.ID).First(user).Error
	if err != nil {
		return err
	}

	return nil
}

// delete user
func (u *userConfig) Delete(user *Order) error {
	err := u.db.Delete(user).Error
	if err != nil {
		return err
	}

	return nil
}

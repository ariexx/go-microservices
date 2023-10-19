package data

import "gorm.io/gorm"

type userConfig struct {
	db *gorm.DB
}

// UserConfig interface
type UserConfig interface {
	AutoMigrate() error
	Create(user *User) error
	Find(user *User) error
	FindByEmail(user *User) error
	FindByID(user *User) error
	Delete(user *User) error
}

func NewUserConfig(db *gorm.DB) UserConfig {
	return &userConfig{db}
}

// User model
type User struct {
	gorm.Model
	Email     string `gorm:"uniqueIndex;not null;type:varchar(100)"`
	Password  string `gorm:"not null"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Role      string `gorm:"not null;sql:enum('admin','user');default:'user'"`
}

// using gorm auto migrate to create tables
func (u *userConfig) AutoMigrate() error {
	err := u.db.AutoMigrate(&User{})
	if err != nil {
		return err
	}

	return nil
}

// using gorm create to insert data
func (u *userConfig) Create(user *User) error {
	err := u.db.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

// using gorm find to get data
func (u *userConfig) Find(user *User) error {
	err := u.db.Find(user).Error
	if err != nil {
		return err
	}

	return nil
}

// using gorm find to get data
func (u *userConfig) FindByEmail(user *User) error {
	err := u.db.Where("email = ?", user.Email).First(user).Error
	if err != nil {
		return err
	}

	return nil
}

// using gorm find to get data
func (u *userConfig) FindByID(user *User) error {
	err := u.db.Where("id = ?", user.ID).First(user).Error
	if err != nil {
		return err
	}

	return nil
}

// delete user
func (u *userConfig) Delete(user *User) error {
	err := u.db.Delete(user).Error
	if err != nil {
		return err
	}

	return nil
}

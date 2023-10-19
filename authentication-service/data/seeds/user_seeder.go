package seeds

import (
	"authentication/data"
	"fmt"
	"gorm.io/gorm"
)

type Seed struct {
	Name string
	Run  func(db *gorm.DB) error
}

type userSeed struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
}

func All() []Seed {
	return []Seed{
		{
			Name: "CreateUser",
			Run: func(db *gorm.DB) error {
				_ = CreateUser(db, userSeed{
					Email:     "arief@gmail.com",
					Password:  "123",
					FirstName: "Arief",
					LastName:  "M",
					Role:      "admin",
				})
				return nil
			},
		},
	}
}

func CreateUser(db *gorm.DB, seed userSeed) error {
	return db.Create(&data.User{
		Email:     seed.Email,
		Password:  seed.Password,
		FirstName: seed.FirstName,
		LastName:  seed.LastName,
		Role:      seed.Role,
	}).Error
}

func Run(db *gorm.DB, seed []Seed) error {
	//looping through all seeds
	for _, s := range seed {
		//run each seed
		err := s.Run(db)
		if err != nil {
			return fmt.Errorf("error while running seed %s: %s", s.Name, err)
		}
	}

	return nil
}

package seeds

import (
	"fmt"
	"gorm.io/gorm"
	"math/rand"
	data "order-service/data"
)

type Seed struct {
	Name string
	Run  func(db *gorm.DB) error
}

type orderSeed data.Order

func All() []Seed {
	return []Seed{
		{
			Name: "CreateOrder",
			Run: func(db *gorm.DB) error {
				_ = CreateOrder(db, orderSeed{
					OrderID:   fmt.Sprintf("ORD-%d", rand.Int()),
					ProductID: fmt.Sprintf("PRD-%d", 1),
					Quantity:  1,
					Price:     1000,
					Total:     1,
				})
				return nil
			},
		},
	}
}

func CreateOrder(db *gorm.DB, seed orderSeed) error {
	err := db.Create(&seed).Error
	if err != nil {
		return err
	}

	return nil
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

package repository

import (
	"context"
	"gorm.io/gorm"
	"order-service/domain/entity"
)

type OrderRepository interface {
	FindOrder(ctx context.Context, id uint) (entity.Order, error)
	CreateOrder(ctx context.Context, order *entity.CreateOrder) (entity.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (o *orderRepository) FindOrder(ctx context.Context, id uint) (entity.Order, error) {
	var order entity.Order
	err := o.db.Find(&order, id).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func (o *orderRepository) CreateOrder(ctx context.Context, order *entity.CreateOrder) (entity.Order, error) {
	var orderEntity entity.Order
	orderEntity = entity.Order{
		Email:     order.Email,
		OrderID:   order.OrderID,
		ProductID: order.ProductID,
		Quantity:  order.Quantity,
		Price:     order.Price,
		Total:     order.Total,
		PlayerID:  order.PlayerID,
	}
	err := o.db.Create(&orderEntity).Error
	if err != nil {
		return orderEntity, err
	}
	return orderEntity, nil
}

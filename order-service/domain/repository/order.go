package repository

import (
	"context"
	"gorm.io/gorm"
	"log"
	"order-service/domain/entity"
)

type OrderRepository interface {
	FindOrder(ctx context.Context, id uint) (entity.Order, error)
	CreateOrder(ctx context.Context, order *entity.CreateOrder) (entity.Order, error)
	FindOrderByOrderID(ctx context.Context, orderID string) (entity.Order, error)
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

func (o *orderRepository) FindOrderByOrderID(ctx context.Context, orderID string) (entity.Order, error) {
	var order entity.Order
	err := o.db.Where("order_id = ?", orderID).First(&order).Error
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
		PaymentID: order.PaymentID,
		Quantity:  order.Quantity,
		Price:     order.Price,
		Total:     order.Total,
		PlayerID:  order.PlayerID,
	}
	err := o.db.Create(&orderEntity).Error
	if err != nil {
		return orderEntity, err
	}
	log.Printf("Email %s\nOrderID %s\nProductID %s\nPaymentID %d\nQuantity %d\nPrice %d\nTotal %d\nPlayerID %s\n", orderEntity.Email, orderEntity.OrderID, orderEntity.ProductID, orderEntity.PaymentID, orderEntity.Quantity, orderEntity.Price, orderEntity.Total, orderEntity.PlayerID)
	return orderEntity, nil
}

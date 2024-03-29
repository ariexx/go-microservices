package service

import (
	"context"
	"order-service/domain/entity"
	"order-service/domain/repository"
)

type OrderService interface {
	CreateOrder(ctx context.Context, order *entity.CreateOrder) (entity.Order, error)
	FindOrder(ctx context.Context, orderId string) (entity.Order, error)
}

type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{
		repo: repo,
	}
}

func (u *orderService) CreateOrder(ctx context.Context, order *entity.CreateOrder) (entity.Order, error) {
	user, err := u.repo.CreateOrder(ctx, order)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *orderService) FindOrder(ctx context.Context, orderId string) (entity.Order, error) {
	user, err := u.repo.FindOrderByOrderID(ctx, orderId)
	if err != nil {
		return user, err
	}
	return user, nil
}

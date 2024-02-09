package gapi

import (
	"context"
	"order-service/domain/repository"
	"order-service/domain/service"
	"order-service/pb"
)

func (s *Server) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (response *pb.GetOrderResponse, err error) {
	//call repository
	orderRepository := repository.NewOrderRepository(s.db)
	orderService := service.NewOrderService(orderRepository)

	order, err := orderService.FindOrder(ctx, req.OrderId)
	if err != nil {
		return nil, err
	}

	response = &pb.GetOrderResponse{
		Order: &pb.Order{
			Email:     order.Email,
			OrderId:   order.OrderID,
			ProductId: order.ProductID,
			PaymentId: int32(order.PaymentID),
			Quantity:  int32(order.Quantity),
			Price:     int32(order.Price),
			Total:     int32(order.Total),
			PlayerId:  order.PlayerID,
		},
	}

	return response, nil
}

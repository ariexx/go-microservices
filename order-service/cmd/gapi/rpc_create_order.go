package gapi

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"math/rand"
	"order-service/data"
	"order-service/domain/entity"
	"order-service/domain/repository"
	"order-service/domain/service"
	"order-service/pb"
	"time"
)

func (s *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	//insert to database
	orderRequest := &data.Order{
		PlayerID:  req.GetPlayerId(),
		Email:     req.GetEmail(),
		OrderID:   fmt.Sprintf("ORD-%d", rand.Intn(1000)),
		ProductID: req.GetProductId(),
		Quantity:  int(req.GetQuantity()),
		Price:     int(req.GetPrice()),
		Total:     int(req.GetTotal()),
	}

	//call repository
	orderRepository := repository.NewOrderRepository(s.db)
	orderService := service.NewOrderService(orderRepository)

	//create order request
	orderEntity := entity.CreateOrder{
		PlayerID:  orderRequest.PlayerID,
		Email:     orderRequest.Email,
		OrderID:   orderRequest.OrderID,
		ProductID: orderRequest.ProductID,
		Quantity:  1,
		Price:     int(orderRequest.Price),
		Total:     int(orderRequest.Total),
	}

	//save order with order service
	createOrder, err := orderService.CreateOrder(ctx, &orderEntity)
	if err != nil {
		return nil, err
	}

	//send an email
	sendEmail(req)

	return &pb.CreateOrderResponse{
		Order: &pb.Order{
			PlayerId:  createOrder.PlayerID,
			Email:     createOrder.Email,
			OrderId:   createOrder.OrderID,
			ProductId: createOrder.ProductID,
			Quantity:  int32(createOrder.Quantity),
			Price:     int32(createOrder.Price),
			Total:     int32(createOrder.Total),
		},
	}, nil
}

func sendEmail(request *pb.CreateOrderRequest) {
	//using goroutine to send email
	go func() {
		client := resty.New().SetTimeout(5 * time.Second)
		_, err := client.R().
			SetHeaders(map[string]string{
				"Content-Type": "application/json",
			}).
			SetBody(map[string]interface{}{
				"email":     request.GetEmail(),
				"player_id": request.GetPlayerId(),
				"product":   request.GetProductId(),
				"price":     request.GetPrice(),
			}).
			Post("http://email-service:5000/v1/email")
		if err != nil {
			fmt.Println("error when calling email service : ", err)
		}
	}()
}

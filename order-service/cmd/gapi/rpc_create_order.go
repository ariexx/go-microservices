package gapi

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
		PaymentID: int(req.GetPaymentId()),
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
		PaymentID: int(orderRequest.PaymentID),
	}

	//save order with order service
	createOrder, err := orderService.CreateOrder(ctx, &orderEntity)
	if err != nil {
		return nil, err
	}

	//send an email
	err = sendEmail(req)
	if err != nil {
		return nil, err
	}

	return &pb.CreateOrderResponse{
		Order: &pb.Order{
			PlayerId:  createOrder.PlayerID,
			Email:     createOrder.Email,
			OrderId:   createOrder.OrderID,
			ProductId: createOrder.ProductID,
			Quantity:  int32(createOrder.Quantity),
			Price:     int32(createOrder.Price),
			Total:     int32(createOrder.Total),
			PaymentId: int32(createOrder.PaymentID),
		},
	}, nil
}

func sendEmail(request *pb.CreateOrderRequest) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) // 3 seconds
	defer cancel()
	//send email using grpc
	conn, err := grpc.DialContext(ctx, "email-service:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Error sending email #1")
		return err
	}

	defer conn.Close()

	client := pb.NewEmailServiceClient(conn)

	_, err = client.SendEmail(ctx, &pb.SendEmailRequest{
		To:           request.GetEmail(),
		PlayerId:     request.GetPlayerId(),
		ProductName:  request.GetProductId(),
		ProductPrice: request.GetPrice(),
	})
	if err != nil {
		fmt.Println("Error sending email #2")
		return err
	}

	return nil
}

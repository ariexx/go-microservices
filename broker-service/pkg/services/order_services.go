package services

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"order-service/pb"
	"time"
)

type OrderService interface {
	CreateOrder(request *pb.CreateOrderRequest) error
}

type orderService struct {
}

func NewOrderServices() OrderService {
	return &orderService{}
}

func (o *orderService) CreateOrder(request *pb.CreateOrderRequest) error {
	conn, err := grpc.Dial("order-service:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Print("Failed to call order-service grpc dial : ", err)
		return err
	}

	defer conn.Close()

	client := pb.NewOrderServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // 3 seconds
	defer cancel()

	_, err = client.CreateOrder(ctx, request)
	if err != nil {
		log.Println("Failed to call order-service grpc CreateOrder : ", err)
	}

	return err
}

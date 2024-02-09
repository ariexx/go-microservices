package services

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"order-service/pb"
	"time"
)

type OrderService interface {
	CreateOrder(request *pb.CreateOrderRequest) (response *pb.CreateOrderResponse, err error)
	GetOrder(request *pb.GetOrderRequest) (response *pb.GetOrderResponse, err error)
}

type orderService struct {
}

func NewOrderServices() OrderService {
	return &orderService{}
}

func (o *orderService) CreateOrder(request *pb.CreateOrderRequest) (response *pb.CreateOrderResponse, err error) {
	conn, err := grpc.Dial("order-service:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Print("Failed to call order-service grpc dial : ", err)
		return nil, fmt.Errorf("%s", "Gagal terhubung ke order-service")
	}

	defer conn.Close()

	client := pb.NewOrderServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // 3 seconds
	defer cancel()

	data, err := client.CreateOrder(ctx, request)
	if err != nil {
		log.Println("Failed to call order-service grpc CreateOrder : ", err)
	}

	return data, err
}

func (o *orderService) GetOrder(request *pb.GetOrderRequest) (response *pb.GetOrderResponse, err error) {
	conn, err := grpc.Dial("order-service:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Print("Failed to call order-service grpc dial : ", err)
		return nil, fmt.Errorf("%s", "Gagal terhubung ke order-service")
	}

	defer conn.Close()

	client := pb.NewOrderServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // 3 seconds
	defer cancel()

	data, err := client.GetOrder(ctx, request)
	if err != nil {
		log.Println("Failed to call order-service grpc GetOrder : ", err)
	}

	return data, err
}

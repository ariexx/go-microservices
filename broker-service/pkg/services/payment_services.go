package services

import (
	"broker_service/pkg/dto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"payment-service/pb"
	"time"
)

type PaymentService interface {
	GetAll() (*[]dto.PaymentResponse, error)
	FindById(id int) (*dto.PaymentResponse, error)
}

type paymentService struct {
}

func NewPaymentService() PaymentService {
	return &paymentService{}
}

func (p *paymentService) GetAll() (*[]dto.PaymentResponse, error) {
	response := new([]dto.PaymentResponse)

	conn, err := grpc.Dial("payment-service:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Failed to call grpc dial : ", err)
		return response, fmt.Errorf("%s", err)
	}

	defer conn.Close()

	client := pb.NewPaymentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) // 3 seconds
	defer cancel()

	res, err := client.GetPayments(ctx, &pb.GetAllEmpty{})
	if err != nil {
		log.Println("Failed to call grpc GetPayments : ", err)
		return response, fmt.Errorf("%s", err)
	}

	var payments []dto.PaymentResponse
	for _, payment := range res.Payments {
		payments = append(payments, dto.PaymentResponse{
			ID:     int64(payment.GetId()),
			Name:   payment.GetName(),
			Banner: payment.GetBanner(),
		})
	}

	return &payments, nil
}

func (p *paymentService) FindById(id int) (*dto.PaymentResponse, error) {
	response := new(dto.PaymentResponse)

	conn, err := grpc.Dial("payment-service:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Failed to call grpc dial : ", err)
		return response, fmt.Errorf("%s", err)
	}

	defer conn.Close()

	client := pb.NewPaymentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) // 3 seconds
	defer cancel()

	res, err := client.GetPayment(ctx, &pb.GetPaymentRequest{
		Id: uint32(id),
	})
	if err != nil {
		log.Println("Failed to call grpc GetPayment : ", err)
		return response, fmt.Errorf("%s", err)
	}

	return &dto.PaymentResponse{
		ID:     int64(res.Payment.GetId()),
		Name:   res.Payment.GetName(),
		Banner: res.Payment.GetBanner(),
	}, nil
}

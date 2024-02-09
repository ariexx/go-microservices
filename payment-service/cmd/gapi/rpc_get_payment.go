package gapi

import (
	"context"
	"log"
	"payment-service/pb"
	"payment-service/pkg/repository"
	"payment-service/pkg/service"
)

func (s *Server) GetPayments(ctx context.Context, req *pb.GetAllEmpty) (*pb.GetAllPaymentResponse, error) {
	//call repository
	paymentRepository := repository.NewPaymentChannelRepository(s.db)
	paymentService := service.NewPaymentService(paymentRepository)

	//get all payment
	payments, err := paymentService.GetAll()
	if err != nil {
		log.Println("Error while getting all payment : ", err)
		return nil, err
	}

	var pbPayment []*pb.Payment

	for _, payment := range payments {
		pbPayment = append(pbPayment, &pb.Payment{
			Id:          uint32(payment.ID),
			Name:        payment.Name,
			Banner:      payment.Banner,
			Description: payment.Description,
		})
	}

	return &pb.GetAllPaymentResponse{
		Payments: pbPayment,
	}, nil

}

func (s *Server) GetPayment(ctx context.Context, req *pb.GetPaymentRequest) (*pb.GetPaymentResponse, error) {
	//call repository
	paymentRepository := repository.NewPaymentChannelRepository(s.db)
	paymentService := service.NewPaymentService(paymentRepository)

	//get payment by id
	payment, err := paymentService.FindByID(int(req.Id))
	if err != nil {
		log.Println("Error while getting payment by id : ", err)
		return nil, err
	}

	return &pb.GetPaymentResponse{
		Payment: &pb.Payment{
			Id:     uint32(payment.ID),
			Name:   payment.Name,
			Banner: payment.Banner,
		},
	}, nil
}

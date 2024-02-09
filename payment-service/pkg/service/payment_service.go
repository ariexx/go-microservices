package service

import (
	"payment-service/pkg/dto"
	"payment-service/pkg/repository"
)

type PaymentService interface {
	GetAll() ([]*dto.PaymentChannelResponse, error)
	FindByID(id int) (*dto.PaymentChannelResponse, error)
}

type paymentService struct {
	paymentRepository repository.PaymentChannelRepository
}

func NewPaymentService(paymentRepository repository.PaymentChannelRepository) PaymentService {
	return &paymentService{paymentRepository: paymentRepository}
}

func (p *paymentService) GetAll() ([]*dto.PaymentChannelResponse, error) {
	paymentChannels, err := p.paymentRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return paymentChannels, nil
}

func (p *paymentService) FindByID(id int) (*dto.PaymentChannelResponse, error) {
	paymentChannel, err := p.paymentRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return paymentChannel, nil
}

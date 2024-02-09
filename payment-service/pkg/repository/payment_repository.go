package repository

import (
	"gorm.io/gorm"
	"payment-service/pkg/dto"
	"payment-service/pkg/model"
)

type PaymentChannelRepository interface {
	FindAll() ([]*dto.PaymentChannelResponse, error)
	FindByID(id int) (*dto.PaymentChannelResponse, error)
}

type paymentChannelRepository struct {
	db *gorm.DB
}

func NewPaymentChannelRepository(db *gorm.DB) PaymentChannelRepository {
	return &paymentChannelRepository{db: db}
}

func (p *paymentChannelRepository) FindAll() ([]*dto.PaymentChannelResponse, error) {
	var paymentChannels []*dto.PaymentChannelResponse
	var paymentModel model.PaymentChannel
	err := p.db.Model(&paymentModel).Find(&paymentChannels).Error
	if err != nil {
		return nil, err
	}
	return paymentChannels, nil
}

func (p *paymentChannelRepository) FindByID(id int) (*dto.PaymentChannelResponse, error) {
	var paymentChannel dto.PaymentChannelResponse
	var paymentModel model.PaymentChannel
	err := p.db.Model(&paymentModel).Where("id = ?", id).First(&paymentChannel).Error
	if err != nil {
		return nil, err
	}
	return &paymentChannel, nil
}

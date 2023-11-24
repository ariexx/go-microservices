package gapi

import (
	"gorm.io/gorm"
	"payment-service/pb"
)

type Server struct {
	pb.UnimplementedPaymentServiceServer
	db *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	return &Server{
		db: db,
	}
}

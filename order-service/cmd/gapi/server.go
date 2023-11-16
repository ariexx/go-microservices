package gapi

import (
	"gorm.io/gorm"
	"order-service/pb"
)

type Server struct {
	pb.UnimplementedOrderServiceServer
	db *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	return &Server{db: db}
}

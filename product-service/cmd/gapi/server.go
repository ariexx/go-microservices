package gapi

import (
	"gorm.io/gorm"
	"product-service/pb"
)

type Server struct {
	pb.UnimplementedProductServiceServer
	db *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	return &Server{
		db: db,
	}
}

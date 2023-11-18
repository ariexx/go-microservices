package gapi

import (
	"context"
	"log"
	"product-service/pb"
	"product-service/pkg/repository"
	"product-service/pkg/service"
)

func (s *Server) GetProducts(ctx context.Context, req *pb.Empty) (*pb.GetAllProductsResponse, error) {
	//call repository
	productRepository := repository.NewProductRepository(s.db)
	productService := service.NewProductService(productRepository)

	//get all products
	products, err := productService.GetAll()
	if err != nil {
		log.Print("Error while getting all products : ", err)
		return nil, err
	}

	return products, nil
}

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

func (s *Server) GetProductById(ctx context.Context, req *pb.GetProductByIdRequest) (*pb.GetProductByIdResponse, error) {
	//call repository
	productRepository := repository.NewProductRepository(s.db)

	//get product by id
	product, err := productRepository.GetProductById(uint(req.GetId()))
	if err != nil {
		log.Print("Error while getting product by id rpc : ", err)
		return nil, err
	}

	newResponse := &pb.GetProductByIdResponse{
		Product: &pb.Product{
			Id:     uint32(product.ID),
			Name:   product.Name,
			Banner: product.Banner,
			ProductDetail: []*pb.ProductDetail{
				{
					Id:        uint32(product.ProductDetails[0].ID),
					ProductId: uint32(product.ProductDetails[0].ProductID),
					Name:      product.ProductDetails[0].Name,
					Price:     product.ProductDetails[0].Price,
				},
			},
		},
	}
	return newResponse, nil
}

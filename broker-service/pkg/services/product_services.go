package services

import (
	"broker_service/pkg/dto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	productProto "product-service/pb"
)

type ProductService interface {
	GetAllProducts() ([]dto.ProductResponse, error)
}

type productService struct {
}

func NewProductServices() ProductService {
	return &productService{}
}

func (p *productService) GetAllProducts() ([]dto.ProductResponse, error) {
	var response []dto.ProductResponse
	conn, err := grpc.Dial("product-service:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Failed to call grpc dial : ", err)
		return response, fmt.Errorf("%s", err)
	}

	defer conn.Close()

	client := productProto.NewProductServiceClient(conn)

	res, err := client.GetProducts(context.Background(), &productProto.Empty{})
	if err != nil {
		log.Println("Failed to call grpc GetProducts : ", err)
		return response, fmt.Errorf("%s", err)
	}

	var products []dto.ProductResponse
	for _, product := range res.Products {
		products = append(products, dto.ProductResponse{
			ID:                    product.GetId(),
			Name:                  product.GetName(),
			Banner:                product.GetBanner(),
			ProductDetailResponse: product.GetProductDetail(),
		})
	}

	return products, nil
}

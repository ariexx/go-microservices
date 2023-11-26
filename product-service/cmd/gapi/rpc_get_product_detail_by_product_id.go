package gapi

import (
	"context"
	"log"
	"product-service/pb"
	"product-service/pkg/repository"
)

func (s *Server) GetProductDetailById(ctx context.Context, req *pb.GetProductDetailByProductIDRequest) (*pb.GetProductDetailByProductIDResponse, error) {
	//call repository
	productRepository := repository.NewProductRepository(s.db)

	productDetailResponse := make([]*pb.ProductDetail, 0)

	//get product detail by product id
	productDetail, err := productRepository.GetByProductID(uint(req.GetProductId()))
	if err != nil {
		log.Print("Error while getting product detail by product id : ", err)
		return nil, err
	}

	for _, product := range productDetail {
		productDetailResponse = append(productDetailResponse, &pb.ProductDetail{
			Id:        uint32(product.ID),
			Name:      product.Name,
			ProductId: uint32(product.ProductID),
			Price:     product.Price,
		})
	}

	return &pb.GetProductDetailByProductIDResponse{
		ProductDetail: productDetailResponse,
	}, nil
}

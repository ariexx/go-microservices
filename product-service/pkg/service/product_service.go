package service

import (
	"fmt"
	"product-service/pb"
	"product-service/pkg/dto"
	"product-service/pkg/repository"
)

type ProductService interface {
	Create(req *dto.CreateProductRequest) error
	GetAll() (*pb.GetAllProductsResponse, error)
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &productService{productRepository: productRepository}
}

func (s *productService) GetAll() (*pb.GetAllProductsResponse, error) {
	products, err := s.productRepository.GetAll()
	if err != nil {
		return nil, err
	}

	fmt.Println("total data : ", len(products))

	var productResponses []*pb.Product

	for _, product := range products {

		//get product detail by product id
		productDetails, err := s.productRepository.GetByProductID(product.ID)
		if err != nil {
			return nil, err
		}

		var productDetailResponses []*pb.ProductDetail
		for _, productDetail := range productDetails {
			productDetailResponses = append(productDetailResponses, &pb.ProductDetail{
				Id:        uint32(productDetail.ID),
				ProductId: uint32(productDetail.ProductID),
				Name:      productDetail.Name,
				Price:     uint32(productDetail.Price),
			})
		}

		productResponses = append(productResponses, &pb.Product{
			Id:            uint32(product.ID),
			Name:          product.Name,
			Banner:        product.Banner,
			ProductDetail: productDetailResponses,
		})

	}

	return &pb.GetAllProductsResponse{
		Products: productResponses,
	}, nil
}

func (s *productService) Create(req *dto.CreateProductRequest) error {
	return s.productRepository.Create(req)
}

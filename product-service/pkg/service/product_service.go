package service

import (
	"product-service/pkg/dto"
	"product-service/pkg/repository"
)

type ProductService interface {
	Create(req *dto.CreateProductRequest) error
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &productService{productRepository: productRepository}
}

func (s *productService) Create(req *dto.CreateProductRequest) error {
	return s.productRepository.Create(req)
}

package repository

import (
	"gorm.io/gorm"
	"log"
	"product-service/pkg/dto"
	"product-service/pkg/model"
)

type ProductRepository interface {
	Create(product *dto.CreateProductRequest) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product *dto.CreateProductRequest) error {
	newProductModel := model.Product{
		Name:   product.ProductName,
		Banner: product.Banner,
	}

	//append product detail
	for _, productDetail := range product.ProductDetails {
		newProductModel.ProductDetails = append(newProductModel.ProductDetails, model.ProductDetail{
			Name:  productDetail.Name,
			Price: productDetail.Price,
		})
	}

	if err := r.db.Create(&newProductModel).Error; err != nil {
		log.Print("Error while creating product : ", err)
		return err
	}

	if err := r.db.Save(&newProductModel).Error; err != nil {
		log.Print("Error while saving product : ", err)
		return err
	}

	return nil
}

package main

import (
	"product-service/config"
	"product-service/pkg/dto"
	"product-service/pkg/repository"
	"product-service/pkg/service"
)

func main() {
	db := config.InitDatabase()

	//call service and repository
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)

	//call create product
	err := productService.Create(&dto.CreateProductRequest{
		ProductName: "Free Fire",
		Banner:      "https://picsum.photos/200/300",
		ProductDetails: []dto.CreateProductDetailRequest{
			{
				Name:  "Diamond 50",
				Price: 5000,
			},
		},
	})
	if err != nil {
		panic("Error while creating product main" + err.Error())
	}

	panic("Success")
}

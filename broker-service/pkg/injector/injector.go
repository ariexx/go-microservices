package injector

import (
	"broker_service/pkg/controller"
	"broker_service/pkg/services"
)

func InitializeProductController() controller.ProductController {
	service := services.NewProductServices()
	productController := controller.NewProductController(service)

	return productController
}
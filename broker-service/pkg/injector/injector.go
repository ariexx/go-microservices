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

func InitializeOrderController() controller.OrderController {
	service := services.NewOrderServices()
	orderController := controller.NewOrderController(service)

	return orderController
}

func InitializePaymentController() controller.PaymentController {
	service := services.NewPaymentService()
	paymentController := controller.NewPaymentController(service)

	return paymentController
}

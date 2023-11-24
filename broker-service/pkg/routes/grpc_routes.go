package routes

import (
	"broker_service/pkg/injector"
	"github.com/gofiber/fiber/v2"
)

func InitializeGRPCRoutes(app *fiber.App) {
	api := app.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			productController := injector.InitializeProductController()
			productController.Routes(v1)

			orderController := injector.InitializeOrderController()
			orderController.Route(v1)
		}
	}
}

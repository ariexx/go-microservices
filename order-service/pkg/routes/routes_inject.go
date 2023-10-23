package routes

import (
	"github.com/gofiber/fiber/v2"
	"order-service/domain/entity"
	"order-service/pkg/helper"
)

func InitializationRoutes(app *fiber.App) {
	app.Post("/order", CreateOrder)
}

func CreateOrder(ctx *fiber.Ctx) error {
	var order entity.CreateOrder
	err := ctx.BodyParser(&order)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ResponseErrorHandler(err.Error()))
	}

}

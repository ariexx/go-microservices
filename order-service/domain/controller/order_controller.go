package controller

import (
	"github.com/gofiber/fiber/v2"
	"order-service/domain/entity"
	"order-service/domain/service"
	"order-service/pkg/helper"
)

type OrderController interface {
	CreateOrder(fiber *fiber.Ctx) error
	Route(fiber fiber.Group, app *fiber.App)
}

type orderController struct {
	service service.OrderService
}

func NewOrderController(service service.OrderService) OrderController {
	return &orderController{
		service: service,
	}
}

func (o *orderController) CreateOrder(ctx *fiber.Ctx) error {
	var order entity.CreateOrder
	err := ctx.BodyParser(&order)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ResponseSuccessHandler(err.Error(), nil))
	}

	orderEntity, err := o.service.CreateOrder(ctx.Context(), &order)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ResponseSuccessHandler(err.Error(), nil))
	}

	return ctx.Status(fiber.StatusOK).JSON(helper.ResponseSuccessHandler("success create order", orderEntity))
}

func (o *orderController) Route(fiber fiber.Group, app *fiber.App) {
	fiber.Post("/order", o.CreateOrder)
}

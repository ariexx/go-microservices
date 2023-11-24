package controller

import (
	"broker_service/pkg/helper"
	"broker_service/pkg/services"
	"github.com/gofiber/fiber/v2"
	"log"
	"order-service/pb"
)

type OrderController interface {
	CreateOrder(ctx *fiber.Ctx) error
	Route(router fiber.Router)
}

type orderController struct {
	orderService services.OrderService
}

func NewOrderController(orderService services.OrderService) OrderController {
	return &orderController{orderService: orderService}
}

func (o *orderController) CreateOrder(ctx *fiber.Ctx) error {
	request := new(pb.CreateOrderRequest)

	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ResponseErrorHandler(err.Error()))
	}

	if err := o.orderService.CreateOrder(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ResponseErrorHandler(err.Error()))
	}

	log.Print("Success create order")
	return ctx.Status(fiber.StatusOK).JSON(helper.ResponseSuccessHandler("success", nil))
}

func (o *orderController) Route(router fiber.Router) {
	router.Post("/orders", o.CreateOrder)
}

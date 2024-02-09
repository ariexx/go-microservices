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
	GetOrder(ctx *fiber.Ctx) error
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

	data, err := o.orderService.CreateOrder(request)
	if err != nil {
		log.Print("Failed create order : ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(helper.ResponseErrorHandler(err.Error()))
	}

	log.Print("Success create order")
	return ctx.Status(fiber.StatusOK).JSON(helper.ResponseSuccessHandler("success", data))
}

func (o *orderController) GetOrder(ctx *fiber.Ctx) error {
	newRequest := &pb.GetOrderRequest{
		OrderId: ctx.Params("id"),
	}

	data, err := o.orderService.GetOrder(newRequest)
	if err != nil {
		log.Print("Failed get order : ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(helper.ResponseErrorHandler(err.Error()))
	}

	log.Print("Success get order")
	return ctx.Status(fiber.StatusOK).JSON(helper.ResponseSuccessHandler("success", data))
}

func (o *orderController) Route(router fiber.Router) {
	router.Post("/orders", o.CreateOrder)
	router.Get("/order/:id", o.GetOrder)
}

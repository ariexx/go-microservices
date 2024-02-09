package controller

import (
	"broker_service/pkg/helper"
	"broker_service/pkg/services"
	"github.com/gofiber/fiber/v2"
	"log"
)

type PaymentController interface {
	GetAll(ctx *fiber.Ctx) error
	GetByID(ctx *fiber.Ctx) error
	Route(app fiber.Router)
}

type paymentController struct {
	paymentService services.PaymentService
}

func NewPaymentController(paymentService services.PaymentService) PaymentController {
	return &paymentController{paymentService: paymentService}
}

func (p *paymentController) Route(app fiber.Router) {
	app.Get("/payments", p.GetAll)
}

func (p *paymentController) GetAll(ctx *fiber.Ctx) error {
	payments, err := p.paymentService.GetAll()
	if err != nil {
		log.Print("Error when calling payment service : ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(helper.ResponseErrorHandler(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(helper.ResponseSuccessHandler("success", payments))
}

type getPaymentById struct {
	id int `params:"id" query:"id"`
}

func (p *paymentController) GetByID(ctx *fiber.Ctx) error {
	id := getPaymentById{}
	if err := ctx.ParamsParser(&id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ResponseErrorHandler(err.Error()))
	}

	payment, err := p.paymentService.FindById(id.id)
	if err != nil {
		log.Print("Error when calling payment service : ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(helper.ResponseErrorHandler(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(helper.ResponseSuccessHandler("success", payment))
}

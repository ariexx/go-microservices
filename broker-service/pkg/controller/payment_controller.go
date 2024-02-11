package controller

import (
	"broker_service/pkg/helper"
	"broker_service/pkg/services"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
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
	app.Get("/payment/:id", p.GetByID)
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
	id := ctx.Params("id")

	idInt, _ := strconv.Atoi(id)

	payment, err := p.paymentService.FindById(idInt)
	if err != nil {
		log.Print("Error when calling payment service : ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(helper.ResponseErrorHandler(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(helper.ResponseSuccessHandler("success", payment))
}

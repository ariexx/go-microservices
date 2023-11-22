package controller

import (
	"broker_service/pkg/helper"
	"broker_service/pkg/services"
	"github.com/gofiber/fiber/v2"
)

type ProductController interface {
	GetAllProducts(ctx *fiber.Ctx) error
	Routes(router fiber.Router)
}

type productController struct {
	productService services.ProductService
}

func (p *productController) GetAllProducts(ctx *fiber.Ctx) error {
	products, err := p.productService.GetAllProducts()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ResponseErrorHandler(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(helper.ResponseSuccessHandler("success", products))
}

func (p *productController) Routes(router fiber.Router) {
	router.Get("/products", p.GetAllProducts)
}

func NewProductController(productService services.ProductService) ProductController {
	return &productController{productService: productService}
}

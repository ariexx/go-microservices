package controller

import (
	"broker_service/pkg/helper"
	"broker_service/pkg/services"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

type ProductController interface {
	GetAllProducts(ctx *fiber.Ctx) error
	Routes(router fiber.Router)
	GetProductDetailByProductId(ctx *fiber.Ctx) error
	GetProductById(ctx *fiber.Ctx) error
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
	router.Get("/products/:productId", p.GetProductDetailByProductId)
	router.Get("/product/:productId", p.GetProductById)
}

func NewProductController(productService services.ProductService) ProductController {
	return &productController{productService: productService}
}

func (p *productController) GetProductDetailByProductId(ctx *fiber.Ctx) error {
	productId := ctx.Params("productId")

	productIdUint, _ := strconv.ParseUint(productId, 10, 32)
	productDetail, err := p.productService.GetProductDetailByProductId(uint32(productIdUint))
	if err != nil {
		log.Println("Error while getting product detail by product id : ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ResponseErrorHandler(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(helper.ResponseSuccessHandler("success", productDetail))
}

func (p *productController) GetProductById(ctx *fiber.Ctx) error {
	productId := ctx.Params("productId")

	productIdUint, _ := strconv.ParseUint(productId, 10, 32)
	product, err := p.productService.GetProductById(uint32(productIdUint))
	if err != nil {
		log.Println("Error while getting product by id : ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ResponseErrorHandler(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(helper.ResponseSuccessHandler("success", product))
}

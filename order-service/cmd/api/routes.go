package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"math/rand"
	"order-service/domain/entity"
	"order-service/domain/repository"
	"order-service/domain/service"
)

type response struct {
	Message string      `json:"message"`
	Error   bool        `json:"error"`
	Data    interface{} `json:"data,omitempty"`
}

type orderRequest struct {
	PlayerID     string `json:"player_id"`
	ProductID    string `json:"product_id"`
	PaymentCode  string `json:"payment_code"`
	ProductPrice int64  `json:"product_price"`
	EmailAccount string `json:"email_account"`
}

func routes(app *fiber.App, db *gorm.DB) {
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(response{
			Message: "pong from order service",
			Error:   false,
		})
	})

	app.Post("/order", func(c *fiber.Ctx) error {
		var orderRequest orderRequest
		if err := c.BodyParser(&orderRequest); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response{
				Message: "Bad Request",
				Error:   true,
			})
		}

		if orderRequest.PlayerID == "" || orderRequest.ProductID == "" || orderRequest.PaymentCode == "" || orderRequest.ProductPrice == 0 || orderRequest.EmailAccount == "" {
			return c.Status(fiber.StatusBadRequest).JSON(response{
				Message: "Invalid Request - Missing Required Fields",
				Error:   true,
			})
		}

		//call repository
		orderRepository := repository.NewOrderRepository(db)
		orderService := service.NewOrderService(orderRepository)

		//create order request
		orderEntity := entity.CreateOrder{
			Email:     orderRequest.EmailAccount,
			OrderID:   fmt.Sprintf("%d", rand.Int()),
			ProductID: orderRequest.ProductID,
			Quantity:  1,
			Price:     int(orderRequest.ProductPrice),
			Total:     int(orderRequest.ProductPrice),
		}

		//save order with order service
		data, err := orderService.CreateOrder(c.Context(), &orderEntity)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response{
				Message: err.Error(),
				Error:   true,
			})
		}

		//TODO send email to user

		return c.Status(fiber.StatusAccepted).JSON(response{
			Message: "Order Success",
			Error:   false,
			Data:    data,
		})
	})
}

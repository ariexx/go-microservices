package main

import "github.com/gofiber/fiber/v2"

type response struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type authPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func routes(app *fiber.App) {
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(response{
			Error:   false,
			Message: "pong from authentication service",
		})
	})

	app.Post("/authenticate", func(c *fiber.Ctx) error {
		var authPayload authPayload
		if err := c.BodyParser(&authPayload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response{
				Error:   true,
				Message: "Bad Request",
			})
		}

		if authPayload.Email == "" || authPayload.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(response{
				Error:   true,
				Message: "Invalid Credential",
			})
		}

		//TODO check credential from database
		if authPayload.Email == "admin@admin.com" && authPayload.Password == "admin" {
			return c.Status(fiber.StatusAccepted).JSON(response{
				Error:   false,
				Message: "Login Success",
				Data:    authPayload,
			})
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(response{
				Error:   true,
				Message: "Invalid Credential",
			})
		}
	})
}

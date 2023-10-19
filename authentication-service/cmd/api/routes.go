package main

import "github.com/gofiber/fiber/v2"

type response struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func routes(app *fiber.App) {
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(response{
			Error:   false,
			Message: "pong from authentication service",
		})
	})
}

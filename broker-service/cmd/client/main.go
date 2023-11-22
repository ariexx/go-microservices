package main

import (
	"broker_service/pkg/routes"
	"errors"
	"github.com/gofiber/fiber/v2"
	"log"
)

const port = ":80"

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			return ctx.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	routes.InitializeGRPCRoutes(app)
	if err := app.Listen(port); err != nil {
		log.Fatalf("Error when starting server: %s", err)
	}

	log.Println("Server broker service started on port : ", port)

}

package main

import (
	"broker_service/cmd/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

const port = ":80"

func main() {

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "*", // "Content-Type,Authorization,Origin,X-Requested-With,Accept",
		ExposeHeaders:    "Link",
		AllowCredentials: true,
		MaxAge:           300,
	}))

	routes.InjectRoutes(app)

	listen := app.Listen(port)
	if listen != nil {
		panic(listen)
	}
}

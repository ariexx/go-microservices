package routes

import (
	"broker_service/cmd/pkg/dto"
	"broker_service/cmd/pkg/helper"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func InjectRoutes(app *fiber.App) {
	app.Get("/ping", ping)
	app.Post("/handle", HandleSubmission)
}

// HandleSubmission implement routes and called to routes function
func HandleSubmission(ctx *fiber.Ctx) error {
	var requestPayload dto.RequestPayload

	if err := ctx.BodyParser(&requestPayload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ResponseSuccessHandler("Bad Request", err))
	}

	switch requestPayload.Action {
	case "auth":
		return authenticate(ctx)
	default:
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ResponseErrorHandler("Bad Request"))
	}
}

func authenticate(ctx *fiber.Ctx) error {
	//var authPayload dto.AuthPayload
	var responseFromAuth helper.ResponseSuccess
	var requestPayload dto.RequestPayload

	if err := ctx.BodyParser(&requestPayload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ResponseSuccessHandler("Bad Request", err))
	}

	//call authenticate service
	client := resty.New()
	resp, err := client.R().SetBody(&requestPayload.Auth).SetResult(&responseFromAuth).Post("http://authentication-service/authenticate")
	if err != nil {
		log.Error("Error when calling authenticate service", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ResponseSuccessHandler("Bad Request", err))
	}

	fmt.Println("response dari auth service : ", resp.String())
	fmt.Println("request ke auth service : ", requestPayload.Auth)

	//check response
	if resp.StatusCode() == fiber.StatusUnauthorized {
		log.Error("Error when calling authenticate service", resp.Status())
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ResponseSuccessHandler("Invalid Credential", resp.Status()))
	} else if resp.StatusCode() != fiber.StatusAccepted && responseFromAuth.Error {
		log.Error("Error when calling authenticate service", resp.Status())
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ResponseSuccessHandler("error calling auth service", resp.Status()))
	}

	//read response from authenticate service
	if responseFromAuth.Error {
		log.Error("Error when calling authenticate service", responseFromAuth.Message)
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ResponseSuccessHandler("Invalid Credential", responseFromAuth.Message))
	}

	return ctx.Status(fiber.StatusAccepted).JSON(helper.ResponseSuccessHandler(fmt.Sprintf("%v", responseFromAuth.Error), responseFromAuth.Message))
}

func ping(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(helper.ResponseSuccessHandler("pong from broker service", nil))
}

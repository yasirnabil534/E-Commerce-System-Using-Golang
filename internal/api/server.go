package api

import (
	"e-commerce-system-using-go/config"
	"log"
	"net/http"
	"github.com/gofiber/fiber/v2"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	app.Get("/health", HealthCheck)

	if err := app.Listen(config.ServerPort); err != nil {
		log.Fatal(err)
	}
}

func HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "OK",
	})
}

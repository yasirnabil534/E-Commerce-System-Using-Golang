package api

import (
	"e-commerce-system-using-golang/config"
	"e-commerce-system-using-golang/internal/api/rest"
	"e-commerce-system-using-golang/internal/api/rest/handlers"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	restHandler := &rest.Resthandler{
		App: app,
	}

	SetupRoutes(restHandler)

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

func SetupRoutes(rh *rest.Resthandler) {
	// User handler
	handlers.SetupUserRoutes(rh)
	// Rest handler
}

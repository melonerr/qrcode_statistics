package routes

import (
	"qrcode_statistics/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	apiGroup := app.Group("/api")
	apiGroup.Get("/health", handlers.Health)
	apiGroup.Get("/users", handlers.CreateMember)
}

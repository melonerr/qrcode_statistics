package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func setupApp() *fiber.App {
	app := fiber.New()
	// Define routes
	app.Post("/users", CreateMember)
	app.Get("/health", Health)
	return app
}

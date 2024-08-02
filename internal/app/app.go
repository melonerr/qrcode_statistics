package app

import (
	"qrcode_statistics/internal/config"
	"qrcode_statistics/internal/middleware"
	"qrcode_statistics/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func Start() {
	config.ConnectDB()

	app := fiber.New()

	middleware.SetupCors(app)

	routes.Setup(app)

	app.Listen(":3000")
}

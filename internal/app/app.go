package app

import (
	"qrcode_statistics/internal/config"
	"qrcode_statistics/internal/middleware"
	"qrcode_statistics/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Start() {
	config.ConnectDB()

	app := fiber.New()

	app.Use(cors.New(middleware.CorsConfig))

	routes.Setup(app)

	app.Listen(":3000")
}

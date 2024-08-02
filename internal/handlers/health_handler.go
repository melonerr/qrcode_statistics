package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Health(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "UP",
		"message": "Service is running",
	})
}

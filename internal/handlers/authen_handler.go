package handlers

import (
	"qrcode_statistics/internal/pkg/models"
	"qrcode_statistics/internal/pkg/repositories"

	"github.com/gofiber/fiber/v2"
)

func Authen(c *fiber.Ctx) error {
	var authen models.Authen
	if err := c.BodyParser(&authen); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	result, err := repositories.Authen(authen)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": result})
}

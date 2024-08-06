package handlers

import (
	"qrcode_statistics/internal/pkg/models"
	"qrcode_statistics/internal/pkg/repositories"

	"github.com/gofiber/fiber/v2"
)

func Statistics(c *fiber.Ctx) error {
	var statistics models.Statistics
	if err := c.BodyParser(&statistics); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	result, err := repositories.AddStatistics(statistics)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

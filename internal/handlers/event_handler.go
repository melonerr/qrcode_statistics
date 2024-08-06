package handlers

import (
	"qrcode_statistics/internal/pkg/models"
	"qrcode_statistics/internal/pkg/repositories"

	"github.com/gofiber/fiber/v2"
)

func GetEventById(c *fiber.Ctx) error {
	id := c.Params("id")
	result, err := repositories.GetEventById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(result)
}

func CreateEvent(c *fiber.Ctx) error {
	var event models.Events
	if err := c.BodyParser(&event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	data := models.Events{
		U_id:       event.U_id,
		Title:      event.Title,
		Detail:     event.Detail,
		Date_start: event.Date_start,
		Date_end:   event.Date_end,
		Status:     true,
	}

	result, err := repositories.CreateEvent(data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

func UpdateEvent(c *fiber.Ctx) error {
	id := c.Params("id")
	var event models.Events
	if err := c.BodyParser(&event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	result, err := repositories.UpdateEvent(id, event)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

func DeleteEvent(c *fiber.Ctx) error {
	id := c.Params("id")
	result, err := repositories.DeleteEvent(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(result)
}

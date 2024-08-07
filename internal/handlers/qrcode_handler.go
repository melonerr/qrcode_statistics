package handlers

import (
	"qrcode_statistics/internal/pkg/models"
	"qrcode_statistics/internal/pkg/repositories"

	"github.com/gofiber/fiber/v2"
)

func GetQrcodeById(c *fiber.Ctx) error {
	id := c.Params("id")
	uid := c.Locals("userID").(string)

	result, err := repositories.GetQrcodeById(id, uid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func CreateQrcode(c *fiber.Ctx) error {
	uid := c.Locals("userID").(string)
	var qrcode models.Qrcode
	if err := c.BodyParser(&qrcode); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	result, err := repositories.CreateQrcode(qrcode, uid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

func UpdateQrcode(c *fiber.Ctx) error {
	id := c.Params("id")
	uid := c.Locals("userID").(string)

	var qrcode models.Qrcode
	if err := c.BodyParser(&qrcode); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	result, err := repositories.UpdateQrcode(id, qrcode, uid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func DeleteQrcode(c *fiber.Ctx) error {
	id := c.Params("id")
	uid := c.Locals("userID").(string)
	result, err := repositories.DeleteQrcode(id, uid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

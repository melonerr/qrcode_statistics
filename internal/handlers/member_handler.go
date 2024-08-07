package handlers

import (
	"qrcode_statistics/internal/pkg/models"
	"qrcode_statistics/internal/pkg/repositories"
	"qrcode_statistics/internal/pkg/service"

	"github.com/gofiber/fiber/v2"
)

// GetMember handles creating a new member
func GetMemberById(c *fiber.Ctx) error {
	id := c.Params("id")
	result, err := repositories.GetMemberById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

// CreateMember handles creating a new member
func CreateMember(c *fiber.Ctx) error {
	var user models.Members
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	data := models.Members{
		Token:    user.Token,
		Username: user.Username,
		Password: service.GenerateMD5Hash(user.Password),
		Email:    user.Email,
		Role:     user.Role,
		Status:   true,
	}

	result, err := repositories.CreateMember(data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

func UpdateMember(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.Members
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	result, err := repositories.UpdateMember(id, user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func DeleteMember(c *fiber.Ctx) error {
	id := c.Params("id")
	result, err := repositories.DeleteMember(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

package handlers

import (
	"qrcode_statistics/internal/models"
	"qrcode_statistics/internal/repositories"

	"github.com/gofiber/fiber/v2"
)

// CreateMember handles creating a new member
func CreateMember(c *fiber.Ctx) error {

	user := models.Members{
		Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
		Username: "melonerr",
		Email:    "melonerr@mail.com",
		Role:     "admin",
	}

	result, err := repositories.CreateMember(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

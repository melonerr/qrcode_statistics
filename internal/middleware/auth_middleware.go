package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *fiber.Ctx) error {
	// Retrieve the JWT token from the context
	user := c.Locals("user").(*jwt.Token)

	// Extract claims from the token
	claims := user.Claims.(jwt.MapClaims)

	// Retrieve the user ID from claims
	id := claims["id"].(string)

	// Store the user ID in the context for use by subsequent handlers
	c.Locals("userID", id)

	// Call the next middleware or handler in the chain
	return c.Next()
}

package routes

import (
	"os"

	"github.com/gofiber/fiber/v3"
)

func Protected() fiber.Handler {
	return func(c fiber.Ctx) error {
		adminKey := os.Getenv("GOMAN_ADMIN_KEY")
		if adminKey == "" {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Admin API key not configured on server",
			})
		}

		clientKey := c.Get("X-API-Key")

		if clientKey != adminKey {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized: Invalid or missing API Key",
			})
		}

		return c.Next()
	}
}

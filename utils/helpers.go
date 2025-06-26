package utils

import "github.com/gofiber/fiber/v2"

func SendJSON(c *fiber.Ctx, status int, data interface{}) error {
	return c.Status(status).JSON(data)
}

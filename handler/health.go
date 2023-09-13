package handler

import "github.com/gofiber/fiber/v2"

// home page
func Health(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "Success!",
		"message": "Hello world!",
		"data":    nil,
	})
}

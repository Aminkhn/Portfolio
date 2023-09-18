package handler

import "github.com/gofiber/fiber/v2"

func GetAllPortfolio(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"name":    "amin",
		"family":  "khalili nasab",
	})
}
func CreatePortfolio(c *fiber.Ctx) error {
	return nil
}
func GetPortfolioById(c *fiber.Ctx) error {
	return nil
}
func GetPortfolioByUsername(c *fiber.Ctx) error {
	return nil
}
func GetPortfolioByEmail(c *fiber.Ctx) error {
	return nil
}
func UpdatePortfolio(c *fiber.Ctx) error {
	return nil
}
func DeletePortfolio(c *fiber.Ctx) error {
	return nil
}

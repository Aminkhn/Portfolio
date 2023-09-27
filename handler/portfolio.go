package handler

import (
	"github.com/aminkhn/portfolio/db/sql"
	"github.com/aminkhn/portfolio/model"
	"github.com/gofiber/fiber/v2"
)

func GetAllPortfolio(c *fiber.Ctx) error {
	// database connection handler
	database, err := sql.MysqlConnect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "failed connecting database!",
			"error":   err.Error(),
		})
	}
	// query handler
	portfolios, err := database.PortfolioQuery.GetAllPortolios() // QUERY
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to retrieve data!",
			"error":   err.Error(),
		})
	}
	// Success result handler
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "success",
		"data":    portfolios,
	})
}
func CreatePortfolio(c *fiber.Ctx) error {
	// database connection handler
	database, err := sql.MysqlConnect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "failed connecting database!",
			"error":   err.Error(),
		})
	}
	// input content check handler
	portfolio := new(model.Portfolio)
	if err := c.BodyParser(portfolio); err != nil {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"status":  "error",
			"message": "there is a problme with your content!",
			"error":   err.Error(),
		})
	}
	// query handler
	err = database.PortfolioQuery.CreatePortfolio(*portfolio) // QUERY
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to retrieve data!",
			"error":   err.Error(),
		})
	}
	// Success result handler
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "success",
	})
}
func GetPortfolioById(c *fiber.Ctx) error {
	// database connection handler
	database, err := sql.MysqlConnect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "failed connecting database!",
			"error":   err.Error(),
		})
	}
	// input content check handler
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"status":  "error",
			"message": "there is a problme with your content!",
			"error":   err.Error(),
		})
	}
	// query handler
	portfolio, err := database.PortfolioQuery.GetPortfolio(id) // QUERY
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to retrieve data!",
			"error":   err.Error(),
		})
	}
	// Success result handler
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "success",
		"data":    portfolio,
	})
}
func GetPortfolioByUsername(c *fiber.Ctx) error {
	return nil
}
func GetPortfolioByEmail(c *fiber.Ctx) error {
	return nil
}
func UpdatePortfolio(c *fiber.Ctx) error {
	// database connection handler
	database, err := sql.MysqlConnect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "failed connecting database!",
			"error":   err.Error(),
		})
	}
	// input content check handler
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"status":  "error",
			"message": "there is a problme with your content!",
			"error":   err.Error(),
		})
	}
	portfolio := new(model.Portfolio)
	if err := c.BodyParser(portfolio); err != nil {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"status":  "error",
			"message": "there is a problme with your content!",
			"error":   err.Error(),
		})
	}
	// query handler
	err = database.PortfolioQuery.UpdatePortfolio(*portfolio, id) // QUERY
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to retrieve data!",
			"error":   err.Error(),
		})
	}
	// Success result handler
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "success",
	})
}
func DeletePortfolio(c *fiber.Ctx) error {
	// database connection handler
	database, err := sql.MysqlConnect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "failed connecting database!",
			"error":   err.Error(),
		})
	}
	// input content check handler
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"status":  "error",
			"message": "there is a problme with your content!",
			"error":   err.Error(),
		})
	}
	portfolio := new(model.Portfolio)
	if err := c.BodyParser(portfolio); err != nil {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"status":  "error",
			"message": "there is a problme with your content!",
			"error":   err.Error(),
		})
	}
	// query handler
	err = database.PortfolioQuery.DeletePortfolio(id) // QUERY
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to retrieve data!",
			"error":   err.Error(),
		})
	}
	// Success result handler
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "success",
	})
}

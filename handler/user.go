package handler

import (
	"github.com/aminkhn/portfolio/db/sql"
	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	database, err := sql.MysqlConnect()
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "failed connecting database!",
			"error":   err,
		})
	}
	users, err := database.UserQuery.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "failed to retrive data!",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    users,
	})
}
func CreateUser(c *fiber.Ctx) error {
	return nil
}
func GetOneUserById(c *fiber.Ctx) error {
	return nil
}
func GetUserByUsername(c *fiber.Ctx) error {
	return nil
}
func GetUserByEmail(c *fiber.Ctx) error {
	return nil
}
func UpdateUser(c *fiber.Ctx) error {
	return nil
}
func DeleteUser(c *fiber.Ctx) error {
	return nil
}

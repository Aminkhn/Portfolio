package handler

import (
	"github.com/aminkhn/portfolio/db/sql"
	"github.com/aminkhn/portfolio/model"
	"github.com/gofiber/fiber/v2"
)

func GetAllUsersHandler(c *fiber.Ctx) error {
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
	users, err := database.UserQuery.GetAllUsers()
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
		"data":    users,
	})
}
func CreateUserHandler(c *fiber.Ctx) error {
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
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"status":  "error",
			"message": "there is a problme with your content!",
			"error":   err.Error(),
		})
	}
	// query handler
	err = database.UserQuery.CreateUser(*user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to retrieve data!",
			"error":   err.Error(),
		})
	}
	// success result handler
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "user created successfuly!",
	})
}
func GetUserHandler(c *fiber.Ctx) error {
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
	user, err := database.UserQuery.GetOneUser(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "failed to execute qurey!",
			"error":   err.Error(),
		})
	}
	// Success result handler
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "success",
		"data":    user,
	})
}
func GetUserByUsername(c *fiber.Ctx) error {
	return nil
}
func GetUserByEmail(c *fiber.Ctx) error {
	return nil
}
func UpdateUserHandler(c *fiber.Ctx) error {
	// database connection handler
	database, err := sql.MysqlConnect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "failed connecting database!",
			"error":   err.Error(),
		})
	}
	// input content handler
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"status":  "error",
			"message": "there is a problme with your content!",
			"error":   "id parameter is empty!",
		})
	}
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"status":  "error",
			"message": "there is a problme with your content!",
			"error":   err.Error(),
		})
	}
	// query handler
	err = database.UserQuery.UpdateUser(*user, id)
	if err != nil {
		return c.Status(fiber.StatusNotModified).JSON(fiber.Map{
			"status":  "error",
			"message": "failed to execute qurey!",
			"error":   err.Error(),
		})
	}
	// success result handler
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "user modified!",
	})
}
func DeleteUserHandler(c *fiber.Ctx) error {
	// database connection handler
	database, err := sql.MysqlConnect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "failed connecting database!",
			"error":   err.Error(),
		})
	}
	// input content handler
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"status":  "error",
			"message": "there is a problme with your content!",
			"error":   "id parameter is empty!",
		})
	}
	// query handler
	if err := database.UserQuery.DeleteUser(id); err != nil {
		return c.Status(fiber.StatusNotModified).JSON(fiber.Map{
			"status":  "error",
			"message": "failed to execute qurey!",
			"error":   err.Error(),
		})
	}
	// success result handler
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "user deleted!",
	})
}

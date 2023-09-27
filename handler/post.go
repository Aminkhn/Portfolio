package handler

import (
	"github.com/aminkhn/portfolio/db/sql"
	"github.com/aminkhn/portfolio/model"
	"github.com/gofiber/fiber/v2"
)

func GetAllPostsHandler(c *fiber.Ctx) error {
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
	posts, err := database.PostQuery.GetAllPosts() // QUERY
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
		"data":    posts,
	})
}

func CreatePostHandler(c *fiber.Ctx) error {
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
	post := new(model.Post)
	if err := c.BodyParser(post); err != nil {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"status":  "error",
			"message": "there is a problme with your content!",
			"error":   err.Error(),
		})
	}
	// query handler
	err = database.PostQuery.CreatePost(*post) // QUERY
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
func GetOnePostByIdHandler(c *fiber.Ctx) error {
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
	post, err := database.PostQuery.GetOnePost(id) // QUERY
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
		"data":    post,
	})
}
func UpdatePostHandler(c *fiber.Ctx) error {
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
	post := new(model.Post)
	if err := c.BodyParser(post); err != nil {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"status":  "error",
			"message": "there is a problme with your content!",
			"error":   err.Error(),
		})
	}
	// query handler
	err = database.PostQuery.UpdatePost(*post, id) // QUERY
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
func DeletePostHandler(c *fiber.Ctx) error {
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
	if err := database.PostQuery.DeletePost(id); err != nil { // QUERY
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

/******************************************/
func GetPublishedPostsHandler(c *fiber.Ctx) error {
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
	posts, err := database.PostQuery.GetPublishedPosts() // QUERY
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
		"data":    posts,
	})
}
func PublishPostHandler(c *fiber.Ctx) error {
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
	post := new(model.Post)
	if err := c.BodyParser(post); err != nil {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"status":  "error",
			"message": "there is a problme with your content!",
			"error":   err.Error(),
		})
	}
	// query handler
	err = database.PostQuery.PublishPost(*post, id) // QUERY
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
		"message": "Post publication successfully changed!",
	})
}

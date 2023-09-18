package router

import (
	"github.com/aminkhn/portfolio/handler"
	"github.com/gofiber/fiber/v2"
)

// user routes
func UserRoute(api fiber.Router) {
	// user group
	u := api.Group("/user")
	// gets all users
	u.Get("/", handler.GetAllUsers)
	// create new user
	u.Post("/", handler.CreateUser)
	// gets one user by id
	u.Get("/:id", handler.GetOneUserById)
	// update user by id
	u.Put("/:id", handler.UpdateUser)
	// edit user by id
	u.Patch("/:id", handler.UpdateUser)
	// delete user by id
	u.Delete("/:id", handler.DeleteUser)
}

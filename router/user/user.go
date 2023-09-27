package user

import (
	"github.com/aminkhn/portfolio/handler"
	"github.com/gofiber/fiber/v2"
)

// user routes
func UserRoute(api fiber.Router) {
	// user group
	u := api.Group("/user")
	// gets all users
	u.Get("/", handler.GetAllUsersHandler)
	// create new user
	u.Post("/", handler.CreateUserHandler)
	// gets one user by id
	u.Get("/:id", handler.GetUserHandler)
	// update user by id
	u.Put("/:id", handler.UpdateUserHandler)
	// edit user by id
	//u.Patch("/:id", handler.UpdateUserHandler)
	// delete user by id
	u.Delete("/:id", handler.DeleteUserHandler)
}

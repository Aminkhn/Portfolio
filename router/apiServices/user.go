package apiServices

import (
	"github.com/aminkhn/portfolio/handler/serviceHandler"
	"github.com/gofiber/fiber/v2"
)

// user routes
func UserRoute(api fiber.Router) {
	// user group
	u := api.Group("/user")
	// gets all users
	u.Get("/", serviceHandler.GetAllUsersHandler)
	// create new user
	u.Post("/", serviceHandler.CreateUserHandler)
	// gets one user by id
	u.Get("/:id", serviceHandler.GetUserHandler)
	// update user by id
	u.Put("/:id", serviceHandler.UpdateUserHandler)
	// edit user by id
	//u.Patch("/:id", serviceHandler.UpdateUserHandler)
	// delete user by id
	u.Delete("/:id", serviceHandler.DeleteUserHandler)
}

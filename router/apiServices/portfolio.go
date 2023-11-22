package apiServices

import (
	"github.com/aminkhn/portfolio/handler/serviceHandler"
	"github.com/gofiber/fiber/v2"
)

func PortfolioRoute(api fiber.Router) {
	// user group
	u := api.Group("/portfolio")
	// gets all users
	u.Get("/", serviceHandler.GetAllPortfolio)
	// create new user
	u.Post("/", serviceHandler.CreatePortfolio)
	// gets one user by id
	u.Get("/:id", serviceHandler.GetPortfolioById)
	// update user by id
	u.Put("/:id", serviceHandler.UpdatePortfolio)
	// edit user by id
	u.Patch("/:id", serviceHandler.UpdatePortfolio)
	// delete user by id
	u.Delete("/:id", serviceHandler.DeletePortfolio)
}

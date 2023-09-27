package portfolio

import (
	"github.com/aminkhn/portfolio/handler"
	"github.com/gofiber/fiber/v2"
)

func PortfolioRoute(api fiber.Router) {
	// user group
	u := api.Group("/portfolio")
	// gets all users
	u.Get("/", handler.GetAllPortfolio)
	// create new user
	u.Post("/", handler.CreatePortfolio)
	// gets one user by id
	u.Get("/:id", handler.GetPortfolioById)
	// update user by id
	u.Put("/:id", handler.UpdatePortfolio)
	// edit user by id
	u.Patch("/:id", handler.UpdatePortfolio)
	// delete user by id
	u.Delete("/:id", handler.DeletePortfolio)
}

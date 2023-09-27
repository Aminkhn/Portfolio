package manage

import "github.com/gofiber/fiber/v2"

func ManageRoute(app fiber.Router) {
	m := app.Group("/cms")
	m.Get("/:id")
	m.Patch("/:id")
}

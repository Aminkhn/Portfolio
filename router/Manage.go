package router

import "github.com/gofiber/fiber/v2"

func ManageRoute(app fiber.Router) {
	m := app.Group("/cms")
	m.Get("/")
	m.Post("/")
	m.Get("/:id")
	m.Put("/:id")
	m.Patch("/:id")
	m.Delete("/:id")
}

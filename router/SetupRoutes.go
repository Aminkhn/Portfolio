package router

import (
	"github.com/aminkhn/portfolio/handler"
	"github.com/gofiber/fiber/v2"
)

// http routes
func SetupRoutes(app *fiber.App) {
	// Health page
	app.Get("/", handler.Health)

}

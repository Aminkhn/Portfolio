package router

import "github.com/gofiber/fiber/v2"

func CMS(app *fiber.App) {
	app.Group("/cms/")
}

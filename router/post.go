package router

import "github.com/gofiber/fiber/v2"

func Post(app *fiber.App) {
	app.Group("/post/")
}

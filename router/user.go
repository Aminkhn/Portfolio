package router

import "github.com/gofiber/fiber/v2"

func User(app *fiber.App) {
	app.Group("/user/")

}

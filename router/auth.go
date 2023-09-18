package router

import (
	"github.com/aminkhn/portfolio/handler"
	"github.com/gofiber/fiber/v2"
)

func AuthRoute(api fiber.Router) {
	a := api.Group("/auth")
	a.Get("/login", handler.Login)
	a.Post("/login", handler.Login)
	a.Get("/logout", handler.Logout)
}

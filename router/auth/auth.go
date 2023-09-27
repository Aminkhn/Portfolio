package auth

import (
	"github.com/aminkhn/portfolio/handler"
	"github.com/gofiber/fiber/v2"
)

func AuthRoute(api fiber.Router) {
	// authentication route group
	a := api.Group("/auth")
	// get login page
	a.Get("/login", handler.Login)
	// handle login op
	a.Post("/login", handler.Login)
	// handle logout op
	a.Get("/logout", handler.Logout)
}

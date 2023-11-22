package apiServices

import (
	"github.com/aminkhn/portfolio/handler/serviceHandler"
	"github.com/gofiber/fiber/v2"
)

func AuthRoute(api fiber.Router) {
	// authentication route group
	a := api.Group("/auth")
	// handle login op
	a.Post("/login", serviceHandler.Login)
	// handle logout op
	a.Get("/logout", serviceHandler.Logout)
}

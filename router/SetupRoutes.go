package router

import (
	"github.com/aminkhn/portfolio/handler"
	"github.com/aminkhn/portfolio/router/auth"
	"github.com/aminkhn/portfolio/router/portfolio"
	"github.com/aminkhn/portfolio/router/post"
	"github.com/aminkhn/portfolio/router/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {

	// General Middlewares
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	// active version
	route_v1(app)
}

// API version 1.0
func route_v1(app *fiber.App) {

	v1 := app.Group("/api/v1")

	// Api Health check
	v1.Get("/", handler.Health)

	// authentication
	auth.AuthRoute(v1)

	// routes
	user.UserRoute(v1)
	post.PostRoute(v1)
	portfolio.PortfolioRoute(v1)
	//manage.ManageRoute(v1)

}

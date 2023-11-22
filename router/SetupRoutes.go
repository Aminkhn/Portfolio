package router

import (
	"github.com/aminkhn/portfolio/handler/serviceHandler"
	"github.com/aminkhn/portfolio/router/apiServices"
	viewpages "github.com/aminkhn/portfolio/router/viewPages"
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

	// routes to Views
	viewRoutes(app)
	// active version
	route_v1(app)
}

// Views
func viewRoutes(app *fiber.App) {
	vP := app.Group("/")
	viewpages.Routes(vP)
}

// API version 1.0
func route_v1(app *fiber.App) {

	v1 := app.Group("/api/v1")

	// Api Health check
	v1.Get("/", serviceHandler.Health)

	// authentication
	apiServices.AuthRoute(v1)

	// routes
	apiServices.UserRoute(v1)
	apiServices.PostRoute(v1)
	apiServices.PortfolioRoute(v1)
	//manage.ManageRoute(v1)

}

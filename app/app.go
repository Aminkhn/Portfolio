package app

import (
	"github.com/aminkhn/portfolio/config"
	"github.com/aminkhn/portfolio/db/cache"
	"github.com/aminkhn/portfolio/router"
	"github.com/gofiber/fiber/v2"
)

func Server() {
	// load env conig
	env, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	// cache (Redis) connection
	cache.RedisConnet()
	// template engine
	//engine := html.New("./views", ".html")
	// new instance of Fiber framework with html engine template
	//app := fiber.New(fiber.Config{Views: engine})
	app := fiber.New()
	// URL routes
	router.SetupRoutes(app)
	// stating webservice
	app.Listen(":" + env.AppPort)

}

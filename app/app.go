package app

import (
	"github.com/aminkhn/portfolio/config"
	"github.com/aminkhn/portfolio/db"
	"github.com/aminkhn/portfolio/router"
	"github.com/gofiber/fiber/v2"
)

func Server() {
	// load env conig
	env, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	// database & cache (Redis) connection for migrations
	db.SetupDatabase()
	// close cache & database connection
	//defer db.DisconnectDatabase()
	// new instance of Fiber framework
	app := fiber.New()
	// URL routes
	router.SetupRoutes(app)
	// stating webservice
	app.Listen(":" + env.AppPort)

}

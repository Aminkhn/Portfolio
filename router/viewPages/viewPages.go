package viewpages

import (
	"github.com/aminkhn/portfolio/handler/pageHandler"
	"github.com/gofiber/fiber/v2"
)

func Routes(app fiber.Router) {
	app.Get("/", pageHandler.Home)
}

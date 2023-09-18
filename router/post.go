package router

import (
	"github.com/aminkhn/portfolio/handler"
	"github.com/gofiber/fiber/v2"
)

func PostRoute(app fiber.Router) {
	p := app.Group("/post")
	p.Get("/", handler.GetAllPosts)
	p.Post("/", handler.CreatePost)
	p.Get("/:id", handler.GetOnePostById)
	p.Put("/:id", handler.UpdatePost)
	p.Patch("/:id", handler.UpdatePost)
	p.Delete("/:id", handler.DeletePost)
}

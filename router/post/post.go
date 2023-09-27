package post

import (
	"github.com/aminkhn/portfolio/handler"
	"github.com/gofiber/fiber/v2"
)

func PostRoute(app fiber.Router) {

	// main post operations
	post := app.Group("/post")

	post.Get("/", handler.GetAllPostsHandler)
	post.Get("/:id", handler.GetOnePostByIdHandler)
	post.Post("/", handler.CreatePostHandler)
	//post.Put("/:id", handler.UpdatePostHandler)
	post.Patch("/:id", handler.UpdatePostHandler)
	post.Delete("/:id", handler.DeletePostHandler)

	// post publish operations
	publish := app.Group("/post_publish")

	publish.Get("/", handler.GetPublishedPostsHandler)
	publish.Patch("/:id", handler.PublishPostHandler)
}

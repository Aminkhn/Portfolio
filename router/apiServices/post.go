package apiServices

import (
	"github.com/aminkhn/portfolio/handler/serviceHandler"
	"github.com/gofiber/fiber/v2"
)

func PostRoute(app fiber.Router) {

	// main post operations
	post := app.Group("/post")

	post.Get("/", serviceHandler.GetAllPostsHandler)
	post.Get("/:id", serviceHandler.GetOnePostByIdHandler)
	post.Post("/", serviceHandler.CreatePostHandler)
	//post.Put("/:id", serviceHandler.UpdatePostHandler)
	post.Patch("/:id", serviceHandler.UpdatePostHandler)
	post.Delete("/:id", serviceHandler.DeletePostHandler)

	// post publish operations
	publish := app.Group("/post_publish")

	publish.Get("/", serviceHandler.GetPublishedPostsHandler)
	publish.Patch("/:id", serviceHandler.PublishPostHandler)
}

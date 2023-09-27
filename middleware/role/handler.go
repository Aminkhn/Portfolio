package role

import "github.com/gofiber/fiber/v2"

func New(Permission string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := roleHandler(c, Permission); err != nil {
			return err
		}
		return c.Next()
	}
}

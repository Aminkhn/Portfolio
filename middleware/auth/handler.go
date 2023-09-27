package auth

import "github.com/gofiber/fiber/v2"

func New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwt, err := NewJWT()
		if err != nil {
			return c.Status(fiber.StatusExpectationFailed).JSON(fiber.Map{
				"message": err,
			})
		}
		token := c.Get("Authorization")

		_, claims, err := jwt.ValidateToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": err,
			})
		}
		c.Locals("claims", claims)
		return c.Next()
	}
}

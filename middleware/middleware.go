package middleware

import (
	"github.com/aminkhn/portfolio/middleware/auth"
	"github.com/aminkhn/portfolio/middleware/role"
	"github.com/gofiber/fiber/v2"
)

func Protected() func(*fiber.Ctx) error {
	return auth.New()
}
func Role(accessType string) func(*fiber.Ctx) error {
	return role.New(accessType)
}

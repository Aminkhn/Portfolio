package role

import (
	"github.com/aminkhn/portfolio/middleware/auth"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrorForbidden = fiber.NewError(fiber.StatusForbidden, "Forbidden")
)
var roles = map[string][]string{
	"super": {"users:read", "users:write", "users:delete", "users:read:all", "users:write:all", "users:delete:all"},
	"admin": {"users:read", "users:write"},
	"user":  {"users:read"},
	"guest": {},
}

func roleHandler(c *fiber.Ctx, permission string) error {
	user := c.Locals("claims").(*auth.CustomClaims)
	// check if user.role in exists in roles
	if roles[user.Role] == nil {
		return ErrorForbidden
	}
	return nil
}

func containsRole(role string, perm string) bool {
	for _, r := range roles[role] {
		if r == perm {
			return true
		}
	}
	return false
}

package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func AuthAdmin() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if !isAdmin(ctx.Get("X-Role")) {
			return ctx.Status(403).JSON(&fiber.Map{
				"message": "Unauthorized",
			})
		}
		return ctx.Next()
	}
}

func isAdmin(role string) bool {
	return role == "Admin"
}

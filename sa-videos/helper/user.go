package helper

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"strconv"
)

func GetUserFromContext(ctx *fiber.Ctx) User {
	userId, err := strconv.ParseUint(ctx.Get("X-User-Id"), 0, 64)
	if err != nil {
		log.Fatal("Fail parse userid", err)
	}

	role := ctx.Get("X-Role")

	if role == "" {
		log.Fatal("Role is empty")
	}

	user := User{
		uint(userId),
		role,
	}
	return user
}

type User struct {
	Id   uint
	Role string
}

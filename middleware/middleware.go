package middleware

import (
	"belajar_golang_fiber/utils"

	"github.com/gofiber/fiber/v2"
)

func Auth(ctx *fiber.Ctx) error {
	token := ctx.Get("x-token")

	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unautenticated",
		})
	}
	_, err := utils.VerifyToken(token)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unautenticated",
		})
	}

// 	if token != "secret" {
// 		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"message": "unautenticated",
// 		})
// 	}
	return ctx.Next()
}


func PermissionCreate(ctx *fiber.Ctx) error {
	return ctx.Next()
}

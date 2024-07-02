package route

import (
	"belajar_golang_fiber/config"
	"belajar_golang_fiber/handler"

	"github.com/gofiber/fiber/v2"
)

func Middelware(ctx *fiber.Ctx) error {
	token := ctx.Get("x-token")

	if token != "secret" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unautenticated",
		})
	}
	return ctx.Next()
}

func RouteInit(r *fiber.App) {
	
	r.Static("/public", config.ProjectRootPath+"/public/asset")
	r.Get("/user",Middelware, handler.UserHandlerGetAll)
	r.Get("/user/:id", handler.UserHandlerGetByID)
	r.Post("/user",handler.UserHandlerCreate)
	r.Put("/user/:id",handler.UserHandlerUpdate)
	r.Delete("/user/:id",handler.UserHandlerDelete)
	
}
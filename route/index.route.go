package route

import (
	"belajar_golang_fiber/config"
	"belajar_golang_fiber/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	
	r.Static("/public", config.ProjectRootPath+"/public/asset")
	r.Get("/user",handler.UserHandlerGetAll)
	r.Get("/user/:id", handler.UserHandlerGetByID)
	r.Post("/user",handler.UserHandlerCreate)
	r.Put("/user/:id",handler.UserHandlerUpdate)
	r.Delete("/user/:id",handler.UserHandlerDelete)
	
}
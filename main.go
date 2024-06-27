package main

import (
	"belajar_golang_fiber/database"
	"belajar_golang_fiber/database/migration"
	"belajar_golang_fiber/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//INITIAL DATABASE
	database.DatabaseInit()
	migration.RunMigration()

	//ROUTE INIT
	route.RouteInit(app)

	app.Listen(":8080")
}
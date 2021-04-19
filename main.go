package main

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-auth-test/db"
	"go-fiber-auth-test/routes"
)

func main() {
	db.GetConnectionDB()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":3000")
}
package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go-fiber-auth-test/handlers"
	"go-fiber-auth-test/middleware"
)

func Setup(app *fiber.App) {
	api := app.Group("/api", logger.New())
	app.Get("/", handlers.Home)

	// auth
	auth := api.Group("/auth")
	auth.Post("/login", handlers.Login)

	// users
	user := api.Group("/user")
	user.Post("/user", middleware.Protected(), handlers.CreateUser)
}

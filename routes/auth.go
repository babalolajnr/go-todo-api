package routes

import (
	"github.com/babalolajnr/go-todo-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func authRoutes(api fiber.Router) {
	auth := api.Group("/auth")

	auth.Post("/login", controllers.Login)
	auth.Post("/register", controllers.Register)
}
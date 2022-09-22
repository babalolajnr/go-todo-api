package routes

import (
	authController "github.com/babalolajnr/go-todo-api/controllers/auth"
	"github.com/gofiber/fiber/v2"
)

func authRoutes(api fiber.Router) {
	auth := api.Group("/auth")

	auth.Post("/login", authController.Login)
	auth.Post("/register", authController.Register)
}
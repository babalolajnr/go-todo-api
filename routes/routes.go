package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func InitializeRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	todoRoutes(api)
	authRoutes(api)
}
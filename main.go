package main

import (
	"log"

	"github.com/babalolajnr/go-todo-api/database"
	"github.com/babalolajnr/go-todo-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	// Establish Database connection
	database.ConnectDB()

	app := fiber.New()

	// Initialize routes
	routes.InitializeRoutes(app)

	app.Use(cors.New())

	app.Get("/health_check", func(c *fiber.Ctx) error {
		return c.SendString("🚀 App is running...")
	})

	log.Fatal(app.Listen(":8000"))
}

package routes

import (
	"github.com/babalolajnr/go-todo-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func todoRoutes(api fiber.Router) {
	todo := api.Group("/todos")

	todo.Get("/", controllers.Todos)
}

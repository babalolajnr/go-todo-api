package routes

import (
	todoController "github.com/babalolajnr/go-todo-api/controllers/todo"
	"github.com/babalolajnr/go-todo-api/middleware"
	"github.com/gofiber/fiber/v2"
)

func todoRoutes(api fiber.Router) {
	todo := api.Group("/todos")

	todo.Get("", middleware.Protected(), todoController.Todos)
	todo.Post("", middleware.Protected(), todoController.Create)
}

package todo

import (
	"github.com/babalolajnr/go-todo-api/controllers/auth"
	"github.com/babalolajnr/go-todo-api/database"
	"github.com/babalolajnr/go-todo-api/models"
	"github.com/gofiber/fiber/v2"
)

func Todos(c *fiber.Ctx) error {
	panic("Not implemented yet")
}

func Create(c *fiber.Ctx) error {
	type CreateTodoInput struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	var createTodoInput CreateTodoInput

	if err := c.BodyParser(&createTodoInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Can't parse request", "data": nil})
	}

	id := auth.ExtractUserId(c)

	todo := models.Todo{
		Title:       createTodoInput.Title,
		Description: createTodoInput.Description,
		Completed:   false,
		UserID:      id,
	}

	db := database.DB.Db

	if err := db.Create(&todo).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't create todo", "data": err})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "created", "message": "Todo Created", "data": todo})
}

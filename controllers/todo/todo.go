package todo

import (
	"github.com/babalolajnr/go-todo-api/controllers/auth"
	"github.com/babalolajnr/go-todo-api/database"
	"github.com/babalolajnr/go-todo-api/models"
	"github.com/gofiber/fiber/v2"
)

func Todos(c *fiber.Ctx) error {

	user, err := auth.AuthenticatedUser(c)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Could not find user", "data": nil})
	}

	var todos []models.Todo

	db := database.DB.Db

	db.Preload("User").Where("user_id = ?", user.ID).Find(&todos)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Todos retrieved", "data": todos})
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

	// Get user
	user, err := auth.GetUserById(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Unable to find user", "data": nil})
	}

	todo := models.Todo{
		Title:       createTodoInput.Title,
		Description: createTodoInput.Description,
		Completed:   false,
	}

	todo.User = *user

	db := database.DB.Db

	if err := db.Create(&todo).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't create todo", "data": err})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "created", "message": "Todo Created", "data": todo})
}

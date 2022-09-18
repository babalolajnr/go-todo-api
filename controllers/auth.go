package controllers

import (
	"errors"
	"time"

	"github.com/babalolajnr/go-todo-api/config"
	"github.com/babalolajnr/go-todo-api/database"
	"github.com/babalolajnr/go-todo-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	type UserData struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var input LoginInput
	var userData UserData

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}

	username := input.Username
	password := input.Password

	user, err := getUserByUsername(username)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid username/password", "data": err})
	}

	if user != nil {
		userData = UserData{
			ID:       user.ID,
			Username: user.Username,
			Password: user.Password,
		}
	}

	if !checkPasswordHash(password, userData.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid username/password", "data": nil})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = userData.Username
	claims["user_id"] = userData.ID
	claims["exp"] = time.Now().Add(time.Hour * 3).Unix()

	t, err := token.SignedString([]byte(config.Config("SECRET")))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}

func getUserByUsername(username string) (*models.User, error) {
	db := database.DB.Db

	var user models.User

	if err := db.Where(&models.User{Username: username}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Register(c *fiber.Ctx) error {
	type RegisterInput struct {
		Name            string `json:"name" form:"name"`
		Username        string `json:"username" form:"username"`
		Password        string `json:"password" form:"password"`
		ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
	}

	type NewUser struct {
		Name     string `json:"name"`
		Username string `json:"username"`
	}

	db := database.DB.Db

	var input RegisterInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}


	if input.Password != input.ConfirmPassword {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": "error", "message": "Password field does not match confirm password field", "data": nil})
	}

	hash, err := hashPassword(input.Password)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})

	}

	user := models.User{
		Name:     input.Name,
		Username: input.Username,
		Password: hash,
	}

	if err := db.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}

	newUser := NewUser{
		Name:     user.Name,
		Username: user.Username,
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": newUser})
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

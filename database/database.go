package database

import (
	"log"

	"github.com/babalolajnr/go-todo-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	Db *gorm.DB
}

var DB DBInstance

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to the database. 😭")
	}

	log.Println("🚀 Database Connected.")

	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("⚗️ Running migrations...")

	db.AutoMigrate(&models.User{}, models.Todo{})

	DB = DBInstance{
		Db: db,
	}
}

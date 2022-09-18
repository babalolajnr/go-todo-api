package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model

	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"not null" json:"description"`
	Completed   bool   `gorm:"not null" json:"completed"`
	UserID      int    `json:"userId"`
	User        User
}
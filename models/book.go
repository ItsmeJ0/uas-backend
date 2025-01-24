package models

import (
	"gorm.io/gorm"
)

type Book struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
	Genre  string `json:"genre"`
}

type Users struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Announcement struct {
	gorm.Model
	Content string `json:"content"`
}

func MigrateSchema(db *gorm.DB) {
	db.AutoMigrate(&Users{}, &Book{}, &Announcement{})
}

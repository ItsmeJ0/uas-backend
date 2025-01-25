package models

import (
	"gorm.io/gorm"
)

type Book struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Year     string `json:"year"`
	Genre    string `json:"genre"`
	Penerbit string `json:"penerbit"`
	Halaman  string `json:"halaman"`
}

type Users struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Announcements struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Message string `json:"message" validate:"required"`
}

func MigrateSchema(db *gorm.DB) {
	db.AutoMigrate(&Users{}, &Book{}, &Announcements{})
}

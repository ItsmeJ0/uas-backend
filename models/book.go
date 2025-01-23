package models

import "gorm.io/gorm"

type Book struct {
	id     uint   `gorm:"primaryKey" json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
	Genre  string `json:"genre"`
}

func MigrateBooks(db *gorm.DB) {
	db.AutoMigrate(&Book{})
}

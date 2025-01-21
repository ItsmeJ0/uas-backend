package models

import "gorm.io/gorm"

type Book struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
	Genre  string `json:"genre"`
}

func MigrateBooks(db *gorm.DB) {
	db.AutoMigrate(&Book{})
}

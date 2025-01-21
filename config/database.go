package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB adalah variabel global untuk koneksi database
var DB *gorm.DB

// ConnectDatabase mengatur koneksi ke database
func ConnectDatabase() {
	// Konfigurasi koneksi
	dsn := "root:@tcp(127.0.0.1:3306)/library?charset=utf8mb4&parseTime=True&loc=Local"
	// Sesuaikan `root`, password, dan nama database (`library`) sesuai dengan kebutuhan Anda.

	// Buka koneksi ke database
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Jika koneksi berhasil, simpan ke variabel global
	DB = database
	log.Println("Database connected successfully!")
}

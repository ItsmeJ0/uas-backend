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
	dsn := "root:ZxPsdaVdKVGwIfczHfHTtuxpUerYHFmy@tcp(autorack.proxy.rlwy.net:50591)/railway?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "admin:admin123@tcp(database-1.cv6oi4oimtxt.ap-southeast-1.rds.amazonaws.com:3306)/books?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "root:@tcp(localhost:3306)/library?charset=utf8&parseTime=True&loc=Local"

	// Buka koneksi ke database
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Jika koneksi berhasil, simpan ke variabel global
	DB = database
	log.Println("Database connected successfully!")
}

// dsn := "root:IXUSUugDwJFcrwkQOzkMXsGgQyBpzDUc@autorack.proxy.rlwy.net:39777/railway"
// mysql://root:IXUSUugDwJFcrwkQOzkMXsGgQyBpzDUc@autorack.proxy.rlwy.net:39777/railway

// dsn := "admin:admin123@tcp(library.cz460gq8ur54.ap-southeast-1.rds.amazonaws.com:3306)/library_db"
// dsn := "root:@tcp(localhost:3306)/library?charset=utf8mb4&parseTime=True&loc=Local"
// Sesuaikan `root`, password, dan nama database (`library`) sesuai dengan kebutuhan Anda.

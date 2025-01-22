package main

import (
	"book-management-backend/config"
	"book-management-backend/models"
	"book-management-backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors" // Import CORS middleware
)

func main() {
	// Initialize Fiber
	app := fiber.New()

	// Connect to Database
	config.ConnectDatabase()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000", // Frontend React berjalan di port 3000
		AllowMethods: "GET,POST,PUT,DELETE",   // Metode HTTP yang diizinkan
		AllowHeaders: "Content-Type",          // Header yang diizinkan
	}))
	// Migrate Database Schema
	models.MigrateBooks(config.DB)

	// Register Routes
	routes.BookRoutes(app)

	// Start Server
	app.Listen(":3001")
}

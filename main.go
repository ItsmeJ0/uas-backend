package main

import (
	"book-management-backend/config"
	"book-management-backend/models"
	"book-management-backend/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize Fiber
	app := fiber.New()

	// Connect to Database
	config.ConnectDatabase()

	// Migrate Database Schema
	models.MigrateBooks(config.DB)

	// Register Routes
	routes.BookRoutes(app)

	// Start Server
	app.Listen(":3000")
}

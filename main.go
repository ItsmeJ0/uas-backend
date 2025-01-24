package main

import (
	"book-management-backend/config"
	"book-management-backend/models"
	"book-management-backend/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors" // Import CORS middleware
	"github.com/gofiber/websocket/v2"             // Import WebSocket package for Fiber
)

var connections []*websocket.Conn

// WebSocket handler
func handleWebSocket(c *websocket.Conn) {
	connections = append(connections, c)

	for {
		_, _, err := c.ReadMessage()
		if err != nil {
			log.Println("WebSocket error:", err)
			break
		}
	}
}

// Function to send announcements to all connected WebSocket clients
func sendAnnouncementToClients(content string) {
	for _, conn := range connections {
		err := conn.WriteMessage(websocket.TextMessage, []byte(content))
		if err != nil {
			log.Println("Error sending message:", err)
			continue
		}
	}
}

func main() {
	// Initialize Fiber
	app := fiber.New()

	// Connect to Database
	config.ConnectDatabase()

	// Use middleware for CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",       // Frontend React runs on port 3000
		AllowMethods: "GET,POST,PUT,DELETE",         // Allowed HTTP methods
		AllowHeaders: "Content-Type, Authorization", // Add 'Authorization' header
	}))

	// Migrate Database Schema
	models.MigrateSchema(config.DB)

	// Register Routes
	routes.BookRoutes(app)

	// WebSocket endpoint for announcements
	app.Get("/ws", websocket.New(handleWebSocket))

	// Route to add announcements
	app.Post("/api/announcement", func(c *fiber.Ctx) error {
		var announcement models.Announcement
		if err := c.BodyParser(&announcement); err != nil {
			return c.Status(400).SendString("Invalid request body")
		}

		// Save announcement to the database
		if err := config.DB.Create(&announcement).Error; err != nil {
			return c.Status(500).SendString("Failed to save announcement")
		}

		// Send the announcement to all connected WebSocket clients
		sendAnnouncementToClients(announcement.Content)

		return c.Status(201).JSON(announcement)
	})

	// Start the server
	app.Listen(":3001")
}

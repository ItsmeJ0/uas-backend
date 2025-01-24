package main

import (
	"book-management-backend/config"
	"book-management-backend/models"
	"book-management-backend/routes"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors" // Import CORS middleware
	"github.com/gofiber/websocket/v2"             // Import WebSocket package for Fiber
)

var connections []*websocket.Conn

// WebSocket handler
func handleWebSocket(c *websocket.Conn) {
	log.Println("New WebSocket connection established") // Log koneksi masuk
	connections = append(connections, c)

	defer func() {
		log.Println("WebSocket connection closed") // Log koneksi tertutup
		for i, conn := range connections {
			if conn == c {
				connections = append(connections[:i], connections[i+1:]...)
				break
			}
		}
	}()

	for {
		_, _, err := c.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)
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
		AllowOrigins:     "https://new-uas-frontend-production.up.railway.app/:3000", // Frontend React runs on port 3000
		AllowMethods:     "GET,POST,PUT,DELETE",                                      // Allowed HTTP methods
		AllowHeaders:     "Content-Type, Authorization",                              // Add 'Authorization' header
		AllowCredentials: true,
	}))

	// Migrate Database Schema
	models.MigrateSchema(config.DB)

	// Register Routes
	routes.BookRoutes(app)

	// WebSocket endpoint for announcements
	app.Get("/ws", websocket.New(handleWebSocket, websocket.Config{
		HandshakeTimeout: 10 * time.Second, // Timeout untuk handshake WebSocket
	}))

	// Route to add announcements
	app.Post("/api/announcement", func(c *fiber.Ctx) error {
		// Cetak body mentah untuk debugging
		log.Println("Raw request body:", string(c.Body()))

		var announcement models.Announcements
		if err := c.BodyParser(&announcement); err != nil {
			log.Println("Error parsing body:", err)
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
		}

		log.Println("Parsed announcement:", announcement)

		if announcement.Message == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Message cannot be empty"})
		}

		if err := config.DB.Create(&announcement).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to save announcement"})
		}

		sendAnnouncementToClients(announcement.Message)
		return c.Status(201).JSON(announcement)
	})

	// Start the server
	app.Listen(":3001")
}

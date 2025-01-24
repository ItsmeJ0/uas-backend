package routes

import (
	"book-management-backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func BookRoutes(app *fiber.App) {
	api := app.Group("/api") // Prefix API
	// Routes for books
	api.Get("/books", handlers.GetBooks)
	api.Post("/books", handlers.CreateBook)
	api.Get("/books/:id", handlers.GetBookByID)
	api.Put("/books/:id", handlers.UpdateBook)
	api.Delete("/books/:id", handlers.DeleteBook)
	// Routes for users
	api.Post("/login", handlers.Login)
	api.Post("/register", handlers.RegisterUser) // Pindahkan ke dalam grup 'api'

}

package routes

import (
	"book-management-backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func BookRoutes(app *fiber.App) {
	api := app.Group("/api") // Prefix API
	api.Get("/books", handlers.GetBooks)
	api.Post("/books", handlers.CreateBook)
	api.Get("/books/:id", handlers.GetBookByID)
	api.Put("/books/:id", handlers.UpdateBook)
	api.Delete("/books/:id", handlers.DeleteBook)
}

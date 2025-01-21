package routes

import (
	"book-management-backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func BookRoutes(app *fiber.App) {
	api := app.Group("/books")
	api.Get("/", handlers.GetBooks)
	api.Get("/:id", handlers.GetBookByID)
	api.Post("/", handlers.CreateBook)
	api.Put("/:id", handlers.UpdateBook)
	api.Delete("/:id", handlers.DeleteBook)
}

package routes

import (
	"book-management-backend/handlers"
	"book-management-backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/login", handlers.Login)
}

func BookRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/books", middlewares.JWTMiddleware, handlers.GetBooks)
	api.Post("/books", middlewares.JWTMiddleware, handlers.CreateBook)
	api.Get("/books/:id", middlewares.JWTMiddleware, handlers.GetBookByID)
	api.Put("/books/:id", middlewares.JWTMiddleware, handlers.UpdateBook)
	api.Delete("/books/:id", middlewares.JWTMiddleware, handlers.DeleteBook)
}

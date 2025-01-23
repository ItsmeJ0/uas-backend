package handlers

import (
	"book-management-backend/config"
	"book-management-backend/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	var books []models.Book
	config.DB.Find(&books)
	return c.JSON(books)
}

func GetBookByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	if result := config.DB.First(&book, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
	}
	return c.JSON(book)
}

func CreateBook(c *fiber.Ctx) error {
	book := new(models.Book)

	// Log body yang diterima
	fmt.Println("Request Body:", string(c.Body()))

	if err := c.BodyParser(book); err != nil {
		fmt.Println("Error parsing body:", err) // Log error parsing
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Log data setelah parsing
	fmt.Printf("Parsed Book: %+v\n", book)

	if result := config.DB.Create(book); result.Error != nil {
		fmt.Println("Error creating book in database:", result.Error) // Log error DB
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create book"})
	}

	return c.Status(201).JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	if result := config.DB.First(&book, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
	}
	if err := c.BodyParser(&book); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	config.DB.Save(&book)
	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	if result := config.DB.First(&book, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
	}
	config.DB.Delete(&book)
	return c.SendString("Book deleted successfully")
}

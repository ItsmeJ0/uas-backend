package handlers

import (
	"book-management-backend/config"
	"book-management-backend/models"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecretKey = []byte("your-secret-key")

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func Login(c *fiber.Ctx) error {
	var request LoginRequest

	// Parse body request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Validasi kredensial (contoh sederhana, gunakan database dalam praktik nyata)
	if request.Username != "admin" || request.Password != "password" {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Buat token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": request.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token berlaku 24 jam
	})

	// Tanda tangani token dengan secret key
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not generate token"})
	}

	return c.JSON(LoginResponse{Token: tokenString})
}
func GetBooks(c *fiber.Ctx) error {
	var books []models.Book
	config.DB.Find(&books)
	fmt.Println(books) // Debug log untuk memastikan ID ada di dalam data
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
	id := c.Params("id") // Ambil ID dari parameter URL
	var book models.Book

	// Cari buku berdasarkan ID
	if result := config.DB.First(&book, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
	}

	// Parse data dari body request
	if err := c.BodyParser(&book); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Update data buku di database
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

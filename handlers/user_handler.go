package handlers

import (
	"book-management-backend/config"
	"book-management-backend/models"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Login(c *fiber.Ctx) error {
	var data struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	fmt.Println("Login data:", data)
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	var user models.Users
	if result := config.DB.Where("email = ?", data.Email).First(&user); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "email not found"})
	}

	// Memeriksa password secara langsung (tanpa hash)
	if user.Password != data.Password {
		return c.Status(400).JSON(fiber.Map{"error": "Incorrect password"})
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret_key"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.JSON(fiber.Map{"token": tokenString})
}

func RegisterUser(c *fiber.Ctx) error {
	user := new(models.Users)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	// Simpan user ke database
	if result := config.DB.Create(user); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.Status(201).JSON(fiber.Map{"message": "User registered successfully"})
}

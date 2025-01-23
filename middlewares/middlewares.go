package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecretKey = []byte("your-secret-key")

func JWTMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")

	// Cek apakah header Authorization ada
	if tokenString == "" {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Parse dan verifikasi token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
		}
		return jwtSecretKey, nil
	})

	// Jika token tidak valid
	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Token valid, lanjutkan ke handler
	return c.Next()
}

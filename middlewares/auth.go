package middlewares

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func AdminMiddleware(c *fiber.Ctx) error {
	godotenv.Load()
	secretKey := os.Getenv("JWT_SECRET")

	tokenString := c.Get("token")
	if tokenString == "" {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{"error": "Token tidak valid"})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return c.Status(401).JSON(fiber.Map{"error": "Token tidak valid"})
	}
	exp := int64(claims["exp"].(float64))
	if time.Now().Unix() > exp {
		return c.Status(401).JSON(fiber.Map{"error": "Token sudah expired"})
	}

	isAdminFloat, ok := claims["is_admin"].(float64)
	if !ok || int(isAdminFloat) != 1 {
		return c.Status(403).JSON(fiber.Map{"error": "Akses ditolak, hanya untuk admin"})
	}

	c.Locals("userID", claims["id"])
	c.Locals("isAdmin", isAdminFloat)

	return c.Next()
}

func AuthMiddleware(c *fiber.Ctx) error {
	godotenv.Load()
	secretKey := os.Getenv("JWT_SECRET")

	tokenString := c.Get("token")
	if tokenString == "" {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{"error": "Token tidak valid"})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return c.Status(401).JSON(fiber.Map{"error": "Token tidak valid"})
	}
	exp := int64(claims["exp"].(float64))
	if time.Now().Unix() > exp {
		return c.Status(401).JSON(fiber.Map{"error": "Token sudah expired"})
	}

	c.Locals("userID", claims["id"])

	return c.Next()
}

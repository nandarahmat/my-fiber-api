package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/nandarahmat/my-fiber-api/database"
	"github.com/nandarahmat/my-fiber-api/models"
	"github.com/nandarahmat/my-fiber-api/routes"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Connect to database
	database.Connect()

	// Auto migrate tables
	database.DB.AutoMigrate(&models.Category{})

	// Initialize Fiber
	app := fiber.New()

	app.Static("/public", "./public")

	// Setup Routes
	routes.SetupRoutes(app)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}

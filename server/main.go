package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/ilham2000r/vue-go-taskmanagment/config"
	"github.com/ilham2000r/vue-go-taskmanagment/routes"
	"log"
)

func main() {
	// Initialize app
	app := fiber.New()

	// Middleware
	app.Use(cors.New())

	// Database connection
	config.ConnectDB()

	// Setup routes
	routes.SetupRoutes(app)

	// Start server
	log.Fatal(app.Listen("localhost:8000"))
}

package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	// Define routes
	setupRoutes(app)

	// Run the server
	log.Fatal(app.Listen(":3000"))
}

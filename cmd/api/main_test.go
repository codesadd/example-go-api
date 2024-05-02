package main

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainT(t *testing.T) {
	// Create a new Fiber app instance
	app := fiber.New()

	// Set up routes
	setupRoutes(app)

	// Perform a GET request to "/items"
	req := httptest.NewRequest(http.MethodGet, "/items", nil)
	resp, err := app.Test(req)

	// Check for errors
	if err != nil {
		t.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

// Add more integration tests as needed for other routes and handlers

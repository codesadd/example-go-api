package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"testing"
)

import (
	"net/http"
	"net/http/httptest"
	"strings"
)

var app *fiber.App

func setup() {
	// Create a new Fiber app instance
	app = fiber.New()

	// Set up routes
	setupRoutes(app)
}

func TestMain(m *testing.M) {
	setup()
	m.Run()
}

func TestGetAllItems(t *testing.T) {
	// Perform a GET request to "/items"
	resp, err := app.Test(httptest.NewRequest(http.MethodGet, "/items", nil))

	// Check for errors
	utils.AssertEqual(t, nil, err)

	// Check the response status code
	utils.AssertEqual(t, http.StatusOK, resp.StatusCode)
}

func TestGetItemByID(t *testing.T) {
	// Perform a GET request to "/items/1"
	resp, err := app.Test(httptest.NewRequest(http.MethodGet, "/items/1", nil))

	// Check for errors
	utils.AssertEqual(t, nil, err)

	// Check the response status code
	utils.AssertEqual(t, http.StatusOK, resp.StatusCode)
}

func TestAddItem(t *testing.T) {
	// Prepare JSON payload for the request
	jsonStr := `{"name":"New Item","price":300}`

	// Perform a POST request to "/items" with the JSON payload
	req := httptest.NewRequest(http.MethodPost, "/items", strings.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	// Check for errors
	utils.AssertEqual(t, nil, err)

	// Check the response status code
	utils.AssertEqual(t, http.StatusCreated, resp.StatusCode)
}

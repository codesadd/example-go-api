package main

import "github.com/gofiber/fiber/v2"

func getAllItems(c *fiber.Ctx) error {
	return c.JSON(items)
}

func getItemByID(c *fiber.Ctx) error {
	itemID, _ := c.ParamsInt("id")

	for _, item := range items {
		if item.ID == itemID {
			return c.JSON(item)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Item not found"})
}

func addItem(c *fiber.Ctx) error {
	var newItem Item
	if err := c.BodyParser(&newItem); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	newItem.ID = len(items) + 1
	items = append(items, newItem)
	return c.Status(fiber.StatusCreated).JSON(newItem)
}

func setupRoutes(app *fiber.App) {
	app.Get("/items", getAllItems)
	app.Get("/items/:id", getItemByID)
	app.Post("/items", addItem)
}

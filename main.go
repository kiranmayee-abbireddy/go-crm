package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Home Route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Go CRM!")
	})

	// Start server
	log.Fatal(app.Listen(":8081"))
}

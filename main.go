package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func setupDatabase() {
	dsn := "host=localhost user=crm_user password=password dbname=crm port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	log.Println("Database connected successfully")
}

func main() {
	app := fiber.New()

	// Setup Database
	setupDatabase()

	// Home Route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Go CRM!")
	})

	// Start server
	log.Fatal(app.Listen(":8081"))
}

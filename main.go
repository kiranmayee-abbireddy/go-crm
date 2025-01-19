package main

import (
	"log"

	"go-crm/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

// Setup the database connection
func setupDatabase() {
	dsn := "host=localhost user=crm_user password=password dbname=crm port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	log.Println("Database connected successfully")

	// Migrate User model to create table
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate models: ", err)
	}
	log.Println("Database migration successful")
}

func main() {
	app := fiber.New()

	// Use logger middleware
	app.Use(logger.New())

	// Setup Database
	setupDatabase()

	// Home Route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Go CRM!")
	})

	// Start server
	log.Fatal(app.Listen(":8081"))
}

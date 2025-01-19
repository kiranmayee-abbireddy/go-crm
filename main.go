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

	// Create a User Route
	app.Post("/users", func(c *fiber.Ctx) error {
		user := new(models.User)

		// Parse the request body into user object
		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Error parsing data")
		}

		// Create user in the database
		result := db.Create(&user)
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error creating user")
		}

		// Return the created user
		return c.Status(fiber.StatusCreated).JSON(user)
	})
	// Fetch all users Route
	app.Get("/users", func(c *fiber.Ctx) error {
		var users []models.User

		// Fetch all users from the database
		if err := db.Find(&users).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error fetching users")
		}

		// Return the users as JSON response
		return c.JSON(users)
	})
	// Fetch user by ID Route
	app.Get("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var user models.User

		// Find user by ID
		if err := db.First(&user, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).SendString("User not found")
		}

		// Return the user as a JSON response
		return c.JSON(user)
	})
	// Update a User Route
	app.Put("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var user models.User

		// Find the user by ID
		if err := db.First(&user, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).SendString("User not found")
		}

		// Parse the updated data
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Error parsing data")
		}

		// Save the updated user data
		if err := db.Save(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error updating user")
		}

		// Return the updated user
		return c.JSON(user)
	})
	// Delete user Route
	app.Delete("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var user models.User

		// Find user by ID
		if err := db.First(&user, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).SendString("User not found")
		}

		// Delete the user
		if err := db.Delete(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error deleting user")
		}

		// Return success response
		return c.SendString("User deleted successfully")
	})

	// Start server
	log.Fatal(app.Listen(":8081"))
}

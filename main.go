package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"

	"github.com/asakew/goLang-Fiber-Postgres/handlers"
	"github.com/asakew/goLang-Fiber-Postgres/models"
)

var DB *gorm.DB

func initDatabase() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&models.Message{})

	DB = db
}

func main() {
	initDatabase()

	app := fiber.New()

	// Serve static files
	app.Static("/", "./static")

	// Register routes
	handlers.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

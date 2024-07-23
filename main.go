package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"os"
)

type Messege struct {
	ID      string `json:"id"`
	Author  string `json:"author"`
	Message string `json:"message"`
}

type Repository struct {
	DB *grom.DB
}

func (r Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/creat_message", r.CreateMessage)
	api.Delete("/delete_message:id", r.deleteMessage)
	api.Get("/get_messages:id", r.getMessagesID)
	api.Get("/messages", r.GetMessages)
}

func (r Repository) CreateMessage(ctx *fiber.Ctx) error {

}

func (r Repository) deleteMessage(ctx *fiber.Ctx) error {

}

func (r Repository) getMessagesID(ctx *fiber.Ctx) error {

}

func (r Repository) GetMessages(ctx *fiber.Ctx) error {

}

func main() {
	godotenv.Load(".env")
	if error != nil {
		panic("Error loading .env file")
	}

	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	
	r := &Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)

}

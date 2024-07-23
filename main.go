package main

import (
	"fmt"
	"github.com/asakew/goLang-Fiber-Postgres/models"
	"github.com/asakew/goLang-Fiber-Postgres/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
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

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/creat_message", r.CreateMessage)
	api.Delete("/delete_message:id", r.deleteMessage)
	api.Get("/get_messages:id", r.getMessagesID)
	api.Get("/messages", r.GetMessages)
}

func (r *Repository) CreateMessage(ctx *fiber.Ctx) error {
	messege := Messege{}

	err := ctx.BodyParser(&messege)
	if err != nil {
		ctx.Status(http.StatusUnprocessableEntity)
		return ctx.JSON(fiber.Map{"status": "error", "message": "Cannot parse JSON", "data": nil})
	}

	err := r.DB.Create(&messege).Error
	if err != nil {
		ctx.Status(http.StatusUnprocessableEntity)
		return ctx.JSON(fiber.Map{"status": "error", "message": "Cannot create message", "data": nil})
	}

	ctx.Status(http.StatusCreated)
	return ctx.JSON(fiber.Map{"status": "success", "message": "Created message", "data": messege})
}

func (r *Repository) deleteMessage(ctx *fiber.Ctx) error {
	Messege.Modal := Models.Messeges{}

	id := ctx.Params("id")
	if id == "" {
		ctx.Status(http.StatusUnprocessableEntity)
		return ctx.JSON(fiber.Map{"status": "error", "message": "Message ID cannot be empty", "data": nil})
	}

	err := r.DB.Delete(&Messege.Modal, id).Error
	if err != nil {
		ctx.Status(http.StatusUnprocessableEntity)
		return ctx.JSON(fiber.Map{"status": "error", "message": "Cannot delete message", "data": nil})
	}

	ctx.Status(http.StatusOK)
	return ctx.JSON(fiber.Map{"status": "success", "message": "Deleted message", "data": nil})
}

func (r *Repository) getMessagesID(ctx *fiber.Ctx) error {
	messegeModels := &[]Models.Messeges{}

	err := r.DB.Find{messegeModels}.Error
	if err != nil {
		ctx.Status(http.StatusUnprocessableEntity)
		return ctx.JSON(fiber.Map{"status": "error", "message": "Cannot find message", "data": nil})
	}

	ctx.Status(http.StatusOK).JSON(fiber.Map{"status": "success", "message": "Fetched all messages", "data": messegeModels})

	return ctx.JSON(messegeModels)
}

func (r *Repository) GetMessages(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	MessegeModal := &ModelsMesseges{}
	if id == "" {
		ctx.Status(http.StatusUnprocessableEntity)
		return ctx.JSON(fiber.Map{"status": "error", "message": "Message ID cannot be empty", "data": nil})
	}

	fmt.Printf("messege id: %s\n", id)
	err := r.DB.Where("id = ?", id).First(&MessegeModal).Error

	err := r.DB.First(&MessegeModal, id).Error
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return ctx.JSON(fiber.Map{"status": "error", "message": "Cannot find message", "data": nil})
	}

	ctx.Status(http.StatusOK)
	return ctx.JSON(fiber.Map{"status": "success", "message": "Fetched all messages", "data": MessegeModal})


}


func main() {
	godotenv.Load(".env") // Load .env file
	if error != nil {
		panic("Error loading .env file")
	}

	//db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL")) // Connect to Postgres
	//if err != nil {
	//	panic(err)
	//}



	Config := storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
	db, err := storage.NewConnection(Config)
	if err != nil {
		panic(err)
	}

	// Connect to Postgres
	db, err = storage.NewConnection(Config{})
	if err != nil {
		panic(err)
	}

	err = models.Messages.MigrateMessages(db)
	if err != nil {
		panic(err)
	}

	r := &Repository{
		DB: db,
	}


	// Custom config
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "goFiber",
		AppName:       "Test App v1.0.1",
	})
	r.SetupRoutes(app)

	app.Static("/static", "./public") // Serve static files

}

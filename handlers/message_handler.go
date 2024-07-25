package handlers

import (
	main "github.com/asakew/goLang-Fiber-Postgres/main"
	"github.com/asakew/goLang-Fiber-Postgres/models"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/messages", GetMessages)
	app.Post("/messages", CreateMessage)
}

func GetMessages(c *fiber.Ctx) error {
	var messages []models.Message
	if result := main.DB.Find(&messages); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}
	return c.JSON(messages)
}

func CreateMessage(c *fiber.Ctx) error {
	msg := new(models.Message)
	if err := c.BodyParser(msg); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}
	if result := main.DB.Create(&msg); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}
	return c.JSON(msg)
}

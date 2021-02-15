package bots

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Bot struct {
}

func NewBot(api fiber.Router, DB *gorm.DB) {

	botRepo := NewBotRepo(DB)

	botService := NewBotService(botRepo)

	botCtrl := NewBotController(botService)

	SetUpRolesRoutes(api, botCtrl)
}

// SetUpRolesRoutes set routes
func SetUpRolesRoutes(api fiber.Router, botCtrl *BotController) {

	api.Get("/bots", botCtrl.Get)
	api.Post("/bots", botCtrl.Create)
	api.Delete("/bots/:id", botCtrl.Delete)

}

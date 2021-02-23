package bots

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Bot struct {
	botRepo       *BotRepo
	botService    *BotService
	botController *BotController
}

func (bot *Bot) NewBot(api fiber.Router, DB *gorm.DB) {

	bot.botRepo = NewBotRepo(DB)

	bot.botService = NewBotService(bot.botRepo)

	bot.botController = NewBotController(bot.botService)

	SetUpRolesRoutes(api, bot.botController)
}

// SetUpRolesRoutes set routes
func SetUpRolesRoutes(api fiber.Router, botCtrl *BotController) {

	api.Get("/bots", botCtrl.Get)
	api.Post("/bots", botCtrl.Create)
	api.Delete("/bots/:id", botCtrl.Delete)

}

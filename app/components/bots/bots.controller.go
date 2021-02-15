package bots

import (
	"net/http"
	"strconv"
	"userManagementApi/app/responses"
	"userManagementApi/app/validation"

	"github.com/gofiber/fiber/v2"
)

type BotController struct {
	botService *BotService
}

func NewBotController(botService *BotService) *BotController {
	return &BotController{
		botService: botService,
	}
}

func (botCtrl *BotController) Create(c *fiber.Ctx) error {
	bot := new(BotsCreateDTO)

	if err := c.BodyParser(bot); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if errors := validation.ValidateStruct(*bot); errors.Errors != nil {
		return responses.SendError(c, errors)
	}

	b, Err := botCtrl.botService.Add(*bot)
	if Err.Errors != nil {
		return responses.SendError(c, Err)
	}
	return responses.Success(c, http.StatusCreated, b)
}

func (botCtrl *BotController) Get(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Query("org"))
	users, Err := botCtrl.botService.Get(id)
	if Err.Errors != nil {
		return responses.SendError(c, Err)
	}
	return responses.Success(c, http.StatusOK, users)
}

func (botCtrl *BotController) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	Err := botCtrl.botService.Delete(id)
	if Err.Errors != nil {
		return responses.SendError(c, Err)
	}
	return c.SendString("Hello, from get roles!")
}

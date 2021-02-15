package bots

import (
	"net/http"
	"userManagementApi/app/responses"
)

type BotService struct {
	botRepo *BotRepo
}

func NewBotService(botRepo *BotRepo) *BotService {
	return &BotService{
		botRepo: botRepo,
	}
}

func (botServ *BotService) Add(bot BotsCreateDTO) (Bots, responses.ErrorData) {
	newBot := Bots{
		Title:           bot.Title,
		OrganizationsID: bot.OrganizationsID,
	}
	b, err := botServ.botRepo.Add(newBot)
	if err != nil {
		return Bots{}, responses.HandleError(http.StatusInternalServerError, err.Error())
	}
	return b, responses.ErrorData{}
}

// Get get all bots
func (botServ *BotService) Get(organizationID int) ([]Bots, responses.ErrorData) {
	bots, err := botServ.botRepo.Get(organizationID)
	if err != nil {
		return []Bots{}, responses.HandleError(http.StatusInternalServerError, err.Error())
	}
	return bots, responses.ErrorData{}
}

// DeleteRole delete bot
func (botServ *BotService) Delete(botID int) responses.ErrorData {
	err := botServ.botRepo.Delete(botID)
	if err != nil {
		return responses.HandleError(http.StatusInternalServerError, err.Error())
	}
	return responses.ErrorData{}
}

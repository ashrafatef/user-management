package bots

import (
	"gorm.io/gorm"
)

type BotRepo struct {
	db *gorm.DB
}

func NewBotRepo(db *gorm.DB) *BotRepo {
	return &BotRepo{
		db: db,
	}
}

// Add add bot
func (botRepo *BotRepo) Add(bot Bots) (Bots, error) {
	res := botRepo.db.Create(&bot)
	if res.Error != nil {
		return Bots{}, res.Error
	}
	return bot, nil
}

// Get get all Bots
func (botRole *BotRepo) Get(organizationID int) ([]Bots, error) {
	var bots []Bots
	res := botRole.db.Raw("SELECT * FROM bots WHERE organizations_id=?", organizationID).Scan(&bots)
	if res.Error != nil {
		return []Bots{}, res.Error
	}
	return bots, nil
}

// Delete bot
func (botRepo *BotRepo) Delete(botID int) error {
	res := botRepo.db.Delete(Bots{
		ID: botID,
	})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

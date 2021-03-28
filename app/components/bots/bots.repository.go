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
	sqlStr := `INSERT INTO blocks_bots (name,slug, bot_id, payload, created_at, updated_at) VALUES ('Welcome Message','welcome-message', ?, '[{"data": {"for": 1600},"block_id": 8}]'::jsonb, current_timestamp, current_timestamp),('Default Message','default-message', ?, '[{"data": {"for": 1600},"block_id": 8}]'::jsonb, current_timestamp, current_timestamp)`
	res = botRepo.db.Exec(sqlStr, bot.ID, bot.ID)
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

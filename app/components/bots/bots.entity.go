package bots

type Bots struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	OrganizationsID int    `json:"organizations_id"`
	BotID           string `json:"bot_id"`
}

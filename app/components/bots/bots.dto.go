package bots

type BotsCreateDTO struct {
	Title           string `json:"title" validate:"required"`
	OrganizationsID int    `json:"organizations_id" validate:"required"`
}

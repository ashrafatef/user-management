package roles

import "time"

type Roles struct {
	ID             int       `json:"id" gorm:"primary_key"`
	Name           string    `json:"name"`
	OrganizationID int       `json:"organization_id"`
	Description    string    `json:"description"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

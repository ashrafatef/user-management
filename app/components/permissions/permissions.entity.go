package permissions

import "time"

type Permissions struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	Name       string    `json: "name"`
	CategoryID uint      `json:"category_id"`
	Type       string    `json:"type"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

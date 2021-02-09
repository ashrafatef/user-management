package users

import "time"

// Organizations_Users model struct
type Organizations_Users struct {
	ID              int       `json:"id" gorm:"primary_key"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	Salt            string    `json:"salt"`
	OrganizationsID int       `json:"organization_id"`
	RoleID          int       `json:"role_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

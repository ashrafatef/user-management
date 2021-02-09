package users

type UserCreateDTO struct {
	FirstName      string `json:"first_name" validate:"required"`
	LastName       string `json:"last_name" validate:"required"`
	Email          string `json:"email" validate:"required"`
	Password       string `json:"password" validate:"required"`
	Salt           string `json:"salt" validate:"required"`
	OrganizationID int    `json:"organization_id" validate:"required"`
	RoleID         int    `json:"role_id" validate:"required"`
}

type UserUpdateDTO struct {
	ID             int    `json:"id" validate:"required"`
	FirstName      string `json:"first_name" validate:"required"`
	LastName       string `json:"last_name" validate:"required"`
	Email          string `json:"email" validate:"required"`
	Salt           string `json:"salt" validate:"required"`
	OrganizationID int    `json:"organization_id" validate:"required"`
	RoleID         int    `json:"role_id" validate:"required"`
}

package roles

type RoleCreateDTO struct {
	Name           string `json:"name" validate:"required,min=3"`
	Description    string `json:"description" validate:""`
	Permissions    []int  `json:"permissions" validate:"required,min=1"`
	OrganizationID int    `json:"organization_id" validate:"required"`
}

type RoleUpdateDTO struct {
	Name           string `json:"name" validate:"required"`
	Description    string `json:"description" validate:""`
	NewAssign      []int  `json:"newAssign" validate:"required"`
	UnAssign       []int  `json:"UnAssign" validate:"required"`
	OrganizationID int    `json:"organization_id" validate:"required"`
}

type RoleDetails struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	PermissionID int    `json:"permission_id"`
}

type SingleRole struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Permissions []int  `json:"permissions"`
}

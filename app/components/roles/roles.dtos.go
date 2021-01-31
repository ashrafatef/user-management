package roles

type RoleCreateDTO struct {
	Name           string `json:"name" validate:"required,min=3"`
	Description    string `json:"description" validate:"required"`
	Permissions    []int  `json:"permissions" validate:"required,min=1"`
	OrganizationID int    `json:"organization_id" validate:"required"`
}

type RoleUpdateDTO struct {
	ID          int    `json:"id" validate:"required"`
	Name        string `json:"name" validate:""`
	Description string `json:"description" validate:""`
	NewAssign   []int  `json:"newAssign" validate:"required"`
	UnAssign    []int  `json:"UnAssign" validate:"required"`
}

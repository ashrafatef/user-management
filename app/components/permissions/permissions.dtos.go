package permissions


type PermissionsCreateDTO struct {
	Name       string `json: "name" validate:"required"`
	CategoryID int    `json:"category_id" validate:"required,oneof=1 2 3 4 5"`
	Type       string `json:"type" validate:"required,oneof=create read update delete"`
}

type PermissionsUpdateDTO struct {
	Name string `json: "name" validate:""`
	ID   int    `json:"id" validate:"required"`
}
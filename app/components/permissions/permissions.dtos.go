package permissions

type PermissionsCreateDTO struct {
	Name     string `json: "name" validate:"required"`
	Category string `json:"category" validate:"required,oneof=builder nlp wordspotting insights settings"`
	Type     string `json:"type" validate:"required,oneof=create read update delete"`
}

type PermissionsUpdateDTO struct {
	Name string `json: "name" validate:""`
	ID   int    `json:"id" validate:"required"`
}

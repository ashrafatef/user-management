package validation

import (
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Failed string
	Tag    string
	Value  string
}

func ValidateStruct(entity interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(entity)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Failed = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

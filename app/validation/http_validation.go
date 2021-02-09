package validation

import (
	"net/http"
	"userManagementApi/app/responses"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Field string
	Tag   string
	Value string
}

func ValidateStruct(entity interface{}) responses.ErrorData {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(entity)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	formattedError := responses.HandleError(http.StatusBadRequest, errors)
	return formattedError
}

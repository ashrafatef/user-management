package error_handler

import (
	"github.com/gofiber/fiber/v2"
)

type ErrorData struct {
	ErrorCode    int         `json:"code"`
	ErrorContent interface{} `json:"error"`
}

type ErrorMeta struct {
	Version string `json:"version"`
}

type ErrorResponseDTO struct {
	Meta ErrorMeta `json:"meta"`
	Data ErrorData `json:"data"`
}

func FormateErrorResponse(statusCode int, errorData error) *ErrorResponseDTO {

	err := ErrorResponseDTO{
		Meta: ErrorMeta{
			Version: "1.5",
		},
		Data: ErrorData{
			ErrorCode:    statusCode,
			ErrorContent: errorData.Error,
		},
	}
	return &err
}

func SendError(ctx *fiber.Ctx, err ErrorResponseDTO) error {
	return ctx.JSON(err)
}

func HandleError(code int, message string) error {

	// log.Fatal("Error: code is %i and message : %s",code, message)
	var msg string = ""
	if len(message) > 0 {
		msg = message
	}
	return fiber.NewError(code, msg)
}

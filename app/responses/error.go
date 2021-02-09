package responses

import (
	"github.com/gofiber/fiber/v2"
)

type ResponseMeta struct {
	Version string `json:"version"`
}

type ErrorData struct {
	Meta   ResponseMeta `json:"meta"`
	Code   int          `json:"code"`
	Errors interface{}  `json:"errors"`
}

func GetResponseMeta() ResponseMeta {
	return ResponseMeta{
		Version: "1.0",
	}
}

func SendError(ctx *fiber.Ctx, err ErrorData) error {
	return ctx.Status(err.Code).JSON(err)
}

func HandleError(code int, errors interface{}) ErrorData {
	err := ErrorData{
		Meta:   GetResponseMeta(),
		Errors: errors,
		Code:   code,
	}
	return err
}

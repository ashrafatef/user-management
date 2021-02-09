package responses

import "github.com/gofiber/fiber/v2"

type SuccessData struct {
	Code int          `json:"code"`
	Meta ResponseMeta `json:"meta"`
	Data interface{}  `json:"data"`
}

func Success(ctx *fiber.Ctx, code int, data interface{}) error {
	success := SuccessData{
		Code: code,
		Meta: GetResponseMeta(),
		Data: data,
	}
	return ctx.Status(code).JSON(success)
}

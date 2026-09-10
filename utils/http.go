package utils

import (
	"github.com/gofiber/fiber/v3"
)

type Resp[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func ErrResponse(c fiber.Ctx, code int, massgae string, data any) error {
	return c.Status(code).JSON(Resp[any]{
		Code:    200,
		Message: massgae,
		Data:    data,
	})
}

func SuccessResp(c fiber.Ctx, data any) error {
	return c.Status(200).JSON(Resp[any]{
		Code:    200,
		Message: "success",
		Data:    data,
	})
}

package main

import (
	"github.com/Edwing123/udem-cine/pkg/codes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func ReadJSONBody[T any](c *fiber.Ctx) (T, error) {
	var v T

	err := c.BodyParser(&v)
	if err != nil {
		return v, err
	}

	return v, nil
}

func (api *Api) GetSession(c *fiber.Ctx) *session.Session {
	return c.Locals(SessionKey).(*session.Session)
}

func (api *Api) ServerError(c *fiber.Ctx, err error) error {
	return SendErrorMessage(c, fiber.StatusInternalServerError, codes.Internal, err)
}

func SendErrorMessage[T any](c *fiber.Ctx, status int, code codes.Code, details T) error {
	return c.Status(status).JSON(ErrorMessage[T]{
		Ok:      false,
		Code:    code,
		Details: details,
	})
}

func SendSucessMessage[T any](c *fiber.Ctx, status int, data T) error {
	return c.Status(status).JSON(SuccessMessage[T]{
		Ok:   true,
		Data: data,
	})
}

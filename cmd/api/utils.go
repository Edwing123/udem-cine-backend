package main

import (
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
	return SendError(c, fiber.StatusInternalServerError, err)
}

func SendError[T any](c *fiber.Ctx, code int, err T) error {
	return c.Status(code).JSON(ErrorMessage[T]{
		Ok:     false,
		Reason: err,
	})
}

func SendMessage[T any](c *fiber.Ctx, code int, data T) error {
	return c.Status(code).JSON(SuccessMessage[T]{
		Ok:   true,
		Data: data,
	})
}

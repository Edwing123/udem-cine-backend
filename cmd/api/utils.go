package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func ReadJSONBody[T any](c *fiber.Ctx) (T, error) {
	var v T

	err := c.BodyParser(&v)
	if err != nil {
		return v, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return v, nil
}

func (api *Api) GetSession(c *fiber.Ctx) *session.Session {
	return c.Locals(SessionKey).(*session.Session)
}

func (api *Api) ServerError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInsufficientStorage).SendString(err.Error())
}

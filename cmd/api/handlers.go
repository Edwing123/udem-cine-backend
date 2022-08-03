package main

import (
	"errors"

	"github.com/Edwing123/udem-cine/pkg/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// Authentication related handlers.
func (api *Api) AuthLogin(c *fiber.Ctx) error {
	credentials, err := ReadJSONBody[models.Credentials](c)
	if err != nil {
		return err
	}

	id, err := api.Models.Authenticate(credentials)
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	sess := c.Locals(SessionKey).(*session.Session)
	sess.Set(UserIdKey, id)
	sess.Set(isLoggedInKey, true)

	return c.SendStatus(fiber.StatusOK)
}

func (api *Api) AuthLogout(c *fiber.Ctx) error {
	sess := c.Locals(SessionKey).(*session.Session)

	err := sess.Destroy()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}

// Users related handlers.
func (api *Api) UserDetaills(c *fiber.Ctx) error {
	sess := api.GetSession(c)
	id := sess.Get(UserIdKey).(int)

	user, err := api.Models.Users.Get(id)
	if errors.Is(err, models.ErroNoRows) {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(user)
}

func (api *Api) UsersList(c *fiber.Ctx) error {
	return c.SendString("Users list")
}

func (api *Api) UsersCreate(c *fiber.Ctx) error {
	return c.SendString("User create")
}

func (api *Api) UsersEdit(c *fiber.Ctx) error {
	return c.SendString("User edit")
}

func (api *Api) UsersDelete(c *fiber.Ctx) error {
	return c.SendString("User delete")
}

// Movies related handlers.
func (api *Api) MoviesList(c *fiber.Ctx) error {
	return c.SendString("Movies list")
}

func (api *Api) MoviesCreate(c *fiber.Ctx) error {
	return c.SendString("Movie create")
}

func (api *Api) MoviesEdit(c *fiber.Ctx) error {
	return c.SendString("Movie edit")
}

func (api *Api) MoviesDelete(c *fiber.Ctx) error {
	return c.SendString("Movie delete")
}

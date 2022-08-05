package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

// Sets store to Fiber context locals.
func (api *Api) SetSessionToContext(c *fiber.Ctx) error {
	sess, err := api.Store.Get(c)
	if err != nil {
		return api.ServerError(c, err)
	}

	c.Locals(SessionKey, sess)

	err = c.Next()
	if err != nil {
		return err
	}

	err = sess.Save()
	if err != nil {
		return api.ServerError(c, err)
	}

	return nil
}

// Calls the next handler only if the user is logged in.
func (api *Api) AuthenticateRequest(c *fiber.Ctx) error {
	sess := api.GetSession(c)
	isLoggedIn, ok := sess.Get(IsLoggedInKey).(bool)

	if !isLoggedIn || !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.Next()
}

func (api *Api) OnlyAdmin(c *fiber.Ctx) error {
	sess := api.GetSession(c)
	id := sess.Get(UserIdKey).(int)

	user, err := api.Models.Users.Get(id)
	if err != nil {
		return api.ServerError(c, err)
	}

	isAdmin := user.Role == "admin"
	if isAdmin {
		return c.Next()
	}

	return SendError(
		c,
		fiber.StatusUnauthorized,
		utils.StatusMessage(fiber.StatusUnauthorized),
	)
}

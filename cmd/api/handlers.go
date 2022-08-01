package main

import "github.com/gofiber/fiber/v2"

// Authentication related handlers.
func (api *Api) AuthLogin(c *fiber.Ctx) error {
	return c.SendString("Login")
}

// Users related handlers.
func (api *Api) UserDetaills(c *fiber.Ctx) error {
	return c.SendString("User details")
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

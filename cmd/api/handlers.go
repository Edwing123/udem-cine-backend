package main

import (
	"errors"
	"fmt"

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
	sess.Set(IsLoggedInKey, true)

	return SendMessage(c, fiber.StatusOK, id)
}

func (api *Api) AuthLogout(c *fiber.Ctx) error {
	sess := c.Locals(SessionKey).(*session.Session)

	err := sess.Destroy()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (api *Api) IsLoggedIn(c *fiber.Ctx) error {
	sess := c.Locals(SessionKey).(*session.Session)
	userId, ok := sess.Get(UserIdKey).(int)
	isok, mm := sess.Get(IsLoggedInKey).(bool)

	fmt.Println(userId, ok)
	fmt.Println(isok, mm)

	return SendMessage(c, fiber.StatusOK, fiber.Map{
		"id": userId,
		"ok": ok,
	})
}

// Users related handlers.
func (api *Api) UserGet(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	user, err := api.Models.Users.Get(id)
	if errors.Is(err, models.ErroNoRows) {
		return SendError(c, fiber.StatusOK, "Usuario no existe")
	}

	return c.JSON(SuccessMessage[models.User]{
		Ok:   true,
		Data: user,
	})
}

func (api *Api) UsersList(c *fiber.Ctx) error {
	users, err := api.Models.Users.List()
	if err != nil {
		return api.ServerError(c, err)
	}

	return SendMessage(c, fiber.StatusOK, users)
}

func (api *Api) UsersCreate(c *fiber.Ctx) error {
	user, err := ReadJSONBody[models.NewUser](c)
	if err != nil {
		return SendError(c, fiber.StatusBadRequest, err)

	}

	err = api.Models.Users.Create(user)
	if err != nil {
		if errors.Is(err, models.ErrUserNameTaken) {
			return SendError(c, fiber.StatusOK, fmt.Sprintf("Usuario %s ya existe", user.Name))
		}

		return api.ServerError(c, err)
	}

	return SendMessage(c, fiber.StatusOK, "Usuario creado")
}

func (api *Api) UsersEdit(c *fiber.Ctx) error {
	user, err := ReadJSONBody[ModelWithId[models.UpdateUser]](c)
	if err != nil {
		return SendError(c, fiber.StatusBadRequest, err)
	}

	err = api.Models.Users.Edit(user.Id, user.Data)
	if err != nil {
		if errors.Is(err, models.ErrUserNameTaken) {
			return SendError(c, fiber.StatusOK, fmt.Sprintf("Usuario %s ya existe", user.Data.Name))
		}

		return api.ServerError(c, err)
	}

	return SendMessage(c, fiber.StatusOK, "Usuario editado")
}

func (api *Api) UsersDelete(c *fiber.Ctx) error {
	body, err := ReadJSONBody[BodyWithId](c)
	if err != nil {
		return SendError(c, fiber.StatusBadRequest, err)
	}

	err = api.Models.Users.Delete(body.Id)
	if err != nil {
		api.ServerError(c, err)
	}

	return SendMessage(c, fiber.StatusOK, "Usuario eliminado")
}

// Movies related handlers.
func (api *Api) MoviesList(c *fiber.Ctx) error {
	return c.SendString("Movies list")
}

func (api *Api) MoviesGet(c *fiber.Ctx) error {
	return c.SendString("Get movie")
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

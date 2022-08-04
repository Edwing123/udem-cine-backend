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

	return SendMessage(c, fiber.StatusOK, fiber.Map{
		"id": userId,
		"ok": ok,
	})
}

// Users related handlers.
func (api *Api) UsersGet(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	user, err := api.Models.Users.Get(id)
	if errors.Is(err, models.ErroNoRows) {
		return SendError(c, fiber.StatusOK, "Usuario no existe")
	}

	return SendMessage(c, fiber.StatusOK, user)
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
	movies, err := api.Models.Movies.List()
	if err != nil {
		return api.ServerError(c, err)
	}

	return SendMessage(c, fiber.StatusOK, movies)
}

func (api *Api) MoviesGet(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	movie, err := api.Models.Movies.Get(id)
	if errors.Is(err, models.ErroNoRows) {
		return SendError(c, fiber.StatusOK, "Pelicula no existe")
	}

	return SendMessage(c, fiber.StatusOK, movie)
}

func (api *Api) MoviesCreate(c *fiber.Ctx) error {
	movie, err := ReadJSONBody[models.NewMovie](c)
	if err != nil {
		return SendError(c, fiber.StatusBadRequest, err)

	}

	err = api.Models.Movies.Create(movie)
	if err != nil {
		if errors.Is(err, models.ErrMovieTitleTaken) {
			return SendError(c, fiber.StatusOK, fmt.Sprintf("Pelicula %s ya existe", movie.Title))
		}

		return api.ServerError(c, err)
	}

	return SendMessage(c, fiber.StatusOK, "Pelicula creada")
}

func (api *Api) MoviesEdit(c *fiber.Ctx) error {
	movie, err := ReadJSONBody[ModelWithId[models.UpdateMovie]](c)
	if err != nil {
		return SendError(c, fiber.StatusBadRequest, err)
	}

	err = api.Models.Movies.Edit(movie.Id, movie.Data)
	if err != nil {
		if errors.Is(err, models.ErrMovieTitleTaken) {
			return SendError(c, fiber.StatusOK, fmt.Sprintf("Pelicula %s ya existe", movie.Data.Title))
		}

		return api.ServerError(c, err)
	}

	return SendMessage(c, fiber.StatusOK, "Pelicula editada")
}

func (api *Api) MoviesDelete(c *fiber.Ctx) error {
	body, err := ReadJSONBody[BodyWithId](c)
	if err != nil {
		return SendError(c, fiber.StatusBadRequest, err)
	}

	err = api.Models.Movies.Delete(body.Id)
	if err != nil {
		api.ServerError(c, err)
	}

	return SendMessage(c, fiber.StatusOK, "Pelicula eliminada")
}

// Rooms related handlers.
func (api *Api) RoomsList(c *fiber.Ctx) error {
	rooms, err := api.Models.Rooms.List()
	if err != nil {
		return api.ServerError(c, err)
	}

	return SendMessage(c, fiber.StatusOK, rooms)
}

func (api *Api) RoomsGet(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("number")
	room, err := api.Models.Rooms.Get(id)
	if errors.Is(err, models.ErroNoRows) {
		return SendError(c, fiber.StatusOK, "Sala no existe")
	}

	return SendMessage(c, fiber.StatusOK, room)
}

func (api *Api) RoomsCreate(c *fiber.Ctx) error {
	room, err := ReadJSONBody[models.NewRoom](c)
	if err != nil {
		return SendError(c, fiber.StatusBadRequest, err)

	}

	err = api.Models.Rooms.Create(room)
	if err != nil {
		if errors.Is(err, models.ErrRoomTaken) {
			return SendError(c, fiber.StatusOK, fmt.Sprintf("Sala %d ya existe", room.Number))
		}

		return api.ServerError(c, err)
	}

	return SendMessage(c, fiber.StatusOK, "Sala creada")
}

func (api *Api) RoomsEdit(c *fiber.Ctx) error {
	room, err := ReadJSONBody[ModelWithId[models.UpdateRoom]](c)
	if err != nil {
		return SendError(c, fiber.StatusBadRequest, err)
	}

	err = api.Models.Rooms.Edit(room.Id, room.Data)
	if err != nil {
		if errors.Is(err, models.ErrRoomTaken) {
			return SendError(c, fiber.StatusOK, fmt.Sprintf("Sala %d ya existe", room.Data.Number))
		}

		return api.ServerError(c, err)
	}

	return SendMessage(c, fiber.StatusOK, "Sala editada")
}

func (api *Api) RoomsDelete(c *fiber.Ctx) error {
	body, err := ReadJSONBody[BodyWithId](c)
	if err != nil {
		return SendError(c, fiber.StatusBadRequest, err)
	}

	err = api.Models.Rooms.Delete(body.Id)
	if err != nil {
		api.ServerError(c, err)
	}

	return SendMessage(c, fiber.StatusOK, "Sala eliminada")
}

// Schedules related handlers.
func (api *Api) SchedulesList(c *fiber.Ctx) error {
	schedules, err := api.Models.Schedules.List()
	if err != nil {
		return api.ServerError(c, err)
	}

	return SendMessage(c, fiber.StatusOK, schedules)
}

func (api *Api) SchedulesGet(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	schedule, err := api.Models.Schedules.Get(id)
	if errors.Is(err, models.ErroNoRows) {
		return SendError(c, fiber.StatusOK, "Horario no existe")
	}

	return SendMessage(c, fiber.StatusOK, schedule)
}

func (api *Api) SchedulesCreate(c *fiber.Ctx) error {
	schedule, err := ReadJSONBody[models.NewSchedule](c)
	if err != nil {
		return SendError(c, fiber.StatusBadRequest, err)

	}

	err = api.Models.Schedules.Create(schedule)
	if err != nil {
		if errors.Is(err, models.ErrScheduleTaken) {
			return SendError(c, fiber.StatusOK, fmt.Sprintf("Horario %v ya existe", schedule.Time))
		}

		return api.ServerError(c, err)
	}

	return SendMessage(c, fiber.StatusOK, "Horari creado")
}

func (api *Api) SchedulesEdit(c *fiber.Ctx) error {
	schedule, err := ReadJSONBody[ModelWithId[models.UpdateSchedule]](c)
	if err != nil {
		return SendError(c, fiber.StatusBadRequest, err)
	}

	err = api.Models.Schedules.Edit(schedule.Id, schedule.Data)
	if err != nil {
		if errors.Is(err, models.ErrScheduleTaken) {
			return SendError(c, fiber.StatusOK, fmt.Sprintf("Horario %v ya existe", schedule.Data.Time))
		}

		return api.ServerError(c, err)
	}

	return SendMessage(c, fiber.StatusOK, "Horario editado")
}

func (api *Api) SchedulesDelete(c *fiber.Ctx) error {
	body, err := ReadJSONBody[BodyWithId](c)
	if err != nil {
		return SendError(c, fiber.StatusBadRequest, err)
	}

	err = api.Models.Schedules.Delete(body.Id)
	if err != nil {
		api.ServerError(c, err)
	}

	return SendMessage(c, fiber.StatusOK, "Horario eliminado")
}

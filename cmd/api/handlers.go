package main

import (
	"errors"
	"fmt"

	"github.com/Edwing123/udem-cine/pkg/codes"
	"github.com/Edwing123/udem-cine/pkg/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// Authentication related handlers.
func (api *Api) AuthLogin(c *fiber.Ctx) error {
	credentials, err := ReadJSONBody[models.Credentials](c)
	if err != nil {
		return SendErrorMessage(c, fiber.StatusBadRequest, codes.Client, err)
	}

	id, err := api.Models.Authenticate(credentials)
	if err != nil {
		if errors.Is(err, codes.AuthFailed) {
			return SendErrorMessage(c, fiber.StatusUnauthorized, codes.AuthFailed, "")
		}

		return api.ServerError(c, err)
	}

	sess := c.Locals(SessionKey).(*session.Session)
	sess.Set(UserIdKey, id)
	sess.Set(IsLoggedInKey, true)

	return SendSucessMessage(c, fiber.StatusOK, fiber.Map{
		"id": id,
	})
}

func (api *Api) AuthLogout(c *fiber.Ctx) error {
	sess := c.Locals(SessionKey).(*session.Session)

	err := sess.Destroy()
	if err != nil {
		return api.ServerError(c, err)
	}

	return SendSucessMessage(c, fiber.StatusOK, "Logout sucessfully")
}

func (api *Api) IsLoggedIn(c *fiber.Ctx) error {
	sess := c.Locals(SessionKey).(*session.Session)
	userId, ok := sess.Get(UserIdKey).(int)

	if !ok {
		return SendErrorMessage(c, fiber.StatusOK, codes.NotLoggedIn, "")
	}

	return SendSucessMessage(c, fiber.StatusOK, fiber.Map{
		"id": userId,
	})
}

// Users related handlers.
func (api *Api) UsersGet(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	user, err := api.Models.Users.Get(id)
	if err != nil {
		if errors.Is(err, codes.NoRecords) {
			return SendErrorMessage(c, fiber.StatusNotFound, codes.NoRecords, "")
		}

		return api.ServerError(c, err)
	}

	return SendSucessMessage(c, fiber.StatusOK, user)
}

func (api *Api) UsersList(c *fiber.Ctx) error {
	users, err := api.Models.Users.List()
	if err != nil {
		return api.ServerError(c, err)
	}

	return SendSucessMessage(c, fiber.StatusOK, users)
}

func (api *Api) UsersCreate(c *fiber.Ctx) error {
	user, err := ReadJSONBody[models.NewUser](c)
	if err != nil {
		return SendErrorMessage(c, fiber.StatusBadRequest, codes.Client, err)

	}

	err = api.Models.Users.Create(user)
	if err != nil {
		if errors.Is(err, codes.UserNameExists) {
			return SendErrorMessage(
				c,
				fiber.StatusConflict,
				codes.UserNameExists,
				fmt.Sprintf("User %s already exists", user.Name),
			)
		}

		return api.ServerError(c, err)
	}

	return SendSucessMessage(c, fiber.StatusCreated, "User created")
}

func (api *Api) UsersEdit(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	user, err := ReadJSONBody[models.UpdateUser](c)
	if err != nil {
		return SendErrorMessage(c, fiber.StatusBadRequest, codes.Client, err)
	}

	err = api.Models.Users.Edit(id, user)
	if err != nil {
		if errors.Is(err, codes.UserNameExists) {
			return SendErrorMessage(
				c,
				fiber.StatusConflict,
				codes.UserNameExists,
				fmt.Sprintf("User %s already exists", user.Name),
			)
		}

		return api.ServerError(c, err)
	}

	return SendSucessMessage(c, fiber.StatusOK, "User edited")
}

func (api *Api) UsersDelete(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	err := api.Models.Users.Delete(id)
	if err != nil {
		api.ServerError(c, err)
	}

	return SendSucessMessage(c, fiber.StatusOK, "User deleted")
}

// Movies related handlers.
func (api *Api) MoviesList(c *fiber.Ctx) error {
	movies, err := api.Models.Movies.List()
	fmt.Println(err)
	if err != nil {
		return api.ServerError(c, err)
	}

	return SendSucessMessage(c, fiber.StatusOK, movies)
}

func (api *Api) MoviesGet(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	movie, err := api.Models.Movies.Get(id)
	if errors.Is(err, codes.NoRecords) {
		return SendErrorMessage(
			c,
			fiber.StatusNotFound,
			codes.NoRecords,
			"",
		)
	}

	return SendSucessMessage(c, fiber.StatusOK, movie)
}

func (api *Api) MoviesCreate(c *fiber.Ctx) error {
	movie, err := ReadJSONBody[models.NewMovie](c)
	if err != nil {
		return SendErrorMessage(c, fiber.StatusBadRequest, codes.Client, err)

	}

	err = api.Models.Movies.Create(movie)
	if err != nil {
		if errors.Is(err, codes.MovieTitleExists) {
			return SendErrorMessage(
				c,
				fiber.StatusConflict,
				codes.MovieTitleExists,
				fmt.Sprintf("Movie %s already exists", movie.Title),
			)
		}

		return api.ServerError(c, err)
	}

	return SendSucessMessage(c, fiber.StatusCreated, "Movie created")
}

func (api *Api) MoviesEdit(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	movie, err := ReadJSONBody[models.UpdateMovie](c)
	if err != nil {
		return SendErrorMessage(c, fiber.StatusBadRequest, codes.Client, err)
	}

	err = api.Models.Movies.Edit(id, movie)
	if err != nil {
		if errors.Is(err, codes.MovieTitleExists) {
			return SendErrorMessage(
				c,
				fiber.StatusConflict,
				codes.MovieTitleExists,
				fmt.Sprintf("Movie %s already exists", movie.Title),
			)
		}

		return api.ServerError(c, err)
	}

	return SendSucessMessage(c, fiber.StatusOK, "Movie edited")
}

func (api *Api) MoviesDelete(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	err := api.Models.Movies.Delete(id)
	if err != nil {
		if errors.Is(err, codes.FunctionDependsOnMovie) {
			return SendErrorMessage(c, fiber.StatusConflict, codes.FunctionDependsOnMovie, "")
		}

		return api.ServerError(c, err)
	}

	return SendSucessMessage(c, fiber.StatusOK, "Movie deleted")
}

// Rooms related handlers.
func (api *Api) RoomsList(c *fiber.Ctx) error {
	rooms, err := api.Models.Rooms.List()
	if err != nil {
		return api.ServerError(c, err)
	}

	return SendSucessMessage(c, fiber.StatusOK, rooms)
}

func (api *Api) RoomsGet(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("number")

	room, err := api.Models.Rooms.Get(id)
	if err != nil {
		if errors.Is(err, codes.NoRecords) {
			return SendErrorMessage(c, fiber.StatusNotFound, codes.NoRecords, "")

		}

		return api.ServerError(c, err)
	}

	return SendSucessMessage(c, fiber.StatusOK, room)
}

func (api *Api) RoomsCreate(c *fiber.Ctx) error {
	room, err := ReadJSONBody[models.NewRoom](c)
	if err != nil {
		return SendErrorMessage(c, fiber.StatusBadRequest, codes.Client, err)
	}

	err = api.Models.Rooms.Create(room)
	if err != nil {
		if errors.Is(err, codes.RoomExists) {
			return SendErrorMessage(
				c,
				fiber.StatusConflict,
				codes.RoomExists,
				fmt.Sprintf("Room %d already exists", room.Number),
			)
		}

		return api.ServerError(c, err)
	}

	return SendSucessMessage(c, fiber.StatusCreated, "Room created")
}

func (api *Api) RoomsEdit(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("number")

	room, err := ReadJSONBody[models.UpdateRoom](c)
	if err != nil {
		return SendErrorMessage(c, fiber.StatusBadRequest, codes.Client, err)
	}

	err = api.Models.Rooms.Edit(id, room)
	if err != nil {
		if errors.Is(err, codes.RoomExists) {
			return SendErrorMessage(
				c,
				fiber.StatusConflict,
				codes.RoomExists,
				fmt.Sprintf("Room %d already exists", room.Number),
			)
		}

		return api.ServerError(c, err)
	}

	return SendSucessMessage(c, fiber.StatusOK, "Room edited")
}

func (api *Api) RoomsDelete(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("number")

	err := api.Models.Rooms.Delete(id)
	if err != nil {
		if errors.Is(err, codes.FunctionDependsOnRoom) {
			return SendErrorMessage(c, fiber.StatusConflict, codes.FunctionDependsOnRoom, "")
		}
		return api.ServerError(c, err)
	}

	return SendSucessMessage(c, fiber.StatusOK, "Room deleted")
}

// Schedules related handlers.
func (api *Api) SchedulesList(c *fiber.Ctx) error {
	schedules, err := api.Models.Schedules.List()
	if err != nil {
		return api.ServerError(c, err)
	}

	return SendSucessMessage(c, fiber.StatusOK, schedules)
}

func (api *Api) SchedulesGet(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	schedule, err := api.Models.Schedules.Get(id)
	if errors.Is(err, codes.NoRecords) {
		return SendErrorMessage(c, fiber.StatusNotFound, codes.NoRecords, "")
	}

	return SendSucessMessage(c, fiber.StatusOK, schedule)
}

func (api *Api) SchedulesCreate(c *fiber.Ctx) error {
	schedule, err := ReadJSONBody[models.NewSchedule](c)
	if err != nil {
		return SendErrorMessage(c, fiber.StatusBadRequest, codes.Client, err)

	}

	err = api.Models.Schedules.Create(schedule)
	if err != nil {
		if errors.Is(err, codes.ScheduleExists) {
			return SendErrorMessage(
				c,
				fiber.StatusConflict,
				codes.ScheduleExists,
				fmt.Sprintf("Schedule %s already exists", schedule.Time),
			)
		}

		return api.ServerError(c, err)
	}

	return SendSucessMessage(c, fiber.StatusCreated, "Schedule created")
}

func (api *Api) SchedulesEdit(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	schedule, err := ReadJSONBody[models.UpdateSchedule](c)
	if err != nil {
		return SendErrorMessage(c, fiber.StatusBadRequest, codes.Client, err)
	}

	err = api.Models.Schedules.Edit(id, schedule)
	if err != nil {
		if errors.Is(err, codes.ScheduleExists) {
			return SendErrorMessage(
				c,
				fiber.StatusConflict,
				codes.ScheduleExists,
				fmt.Sprintf("Schedule %s already exists", schedule.Time),
			)
		}

		return api.ServerError(c, err)
	}

	return SendSucessMessage(c, fiber.StatusOK, "Schedule edited")
}

func (api *Api) SchedulesDelete(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	err := api.Models.Schedules.Delete(id)
	if err != nil {
		if errors.Is(err, codes.FunctionDependsOnSchedule) {
			return SendErrorMessage(c, fiber.StatusConflict, codes.FunctionDependsOnSchedule, "")
		}

		return api.ServerError(c, err)
	}

	return SendSucessMessage(c, fiber.StatusOK, "schedule deleted")
}

// Functions related handlers.
func (api *Api) FunctionsList(c *fiber.Ctx) error {
	functions, err := api.Models.Functions.List()
	if err != nil {
		return api.ServerError(c, err)
	}

	return SendSucessMessage(c, fiber.StatusOK, functions)
}

func (api *Api) FunctionsGet(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	function, err := api.Models.Functions.Get(id)
	if errors.Is(err, codes.NoRecords) {
		return SendErrorMessage(c, fiber.StatusNotFound, codes.NoRecords, "")
	}

	return SendSucessMessage(c, fiber.StatusOK, function)
}

func (api *Api) FunctionsCreate(c *fiber.Ctx) error {
	function, err := ReadJSONBody[models.NewFunction](c)
	if err != nil {
		return SendErrorMessage(c, fiber.StatusBadRequest, codes.Client, err)

	}

	err = api.Models.Functions.Create(function)
	if err != nil {
		if errors.Is(err, codes.FunctionRoomScheduleConflict) {
			return SendErrorMessage(
				c,
				fiber.StatusConflict,
				codes.FunctionRoomScheduleConflict,
				"Function (schedule, room) pair are in conflict with another function",
			)
		}

		return api.ServerError(c, err)
	}

	return SendSucessMessage(c, fiber.StatusCreated, "Function created")
}

func (api *Api) FunctionsEdit(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("number")

	function, err := ReadJSONBody[models.UpdateFunction](c)
	if err != nil {
		return SendErrorMessage(c, fiber.StatusBadRequest, codes.Client, err)
	}

	err = api.Models.Functions.Edit(id, function)
	if err != nil {
		if errors.Is(err, codes.FunctionRoomScheduleConflict) {
			return SendErrorMessage(
				c,
				fiber.StatusConflict,
				codes.RoomExists,
				"Function (schedule, room) pair are in conflict with another function",
			)
		}

		return api.ServerError(c, err)
	}

	return SendSucessMessage(c, fiber.StatusOK, "Function edited")
}

func (api *Api) FunctionsDelete(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	err := api.Models.Functions.Archive(id)
	if err != nil {
		if errors.Is(err, codes.FunctionRoomScheduleConflict) {
			return SendErrorMessage(c, fiber.StatusConflict, codes.FunctionRoomScheduleConflict, "")
		}

		return api.ServerError(c, err)
	}

	return SendSucessMessage(c, fiber.StatusOK, "Function deleted")
}

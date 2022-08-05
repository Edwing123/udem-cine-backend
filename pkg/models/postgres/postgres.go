package postgres

import (
	"context"

	"github.com/Edwing123/udem-cine/pkg/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

func New(dsn string) (models.Models, error) {
	ctx := context.Background()

	conn, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return models.Models{}, err
	}

	if err := conn.Ping(ctx); err != nil {
		return models.Models{}, err
	}

	auth := &AuthController{
		conn,
	}

	movies := &MoviesController{
		conn,
	}

	users := &UsersController{
		conn,
	}

	rooms := &RoomsController{
		conn,
	}

	schedules := &SchedulesController{
		conn,
	}

	functions := &FunctionsController{
		conn,
	}

	return models.Models{
		Auth:      auth,
		Movies:    movies,
		Users:     users,
		Rooms:     rooms,
		Schedules: schedules,
		Functions: functions,
	}, nil
}

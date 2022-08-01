package postgres

import (
	"context"

	"github.com/Edwing123/udem-cine/pkg/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Config struct{}

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

	return models.Models{
		Auth:   auth,
		Movies: movies,
		Users:  users,
	}, nil
}

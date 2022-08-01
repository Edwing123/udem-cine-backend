package postgres

import (
	"github.com/Edwing123/udem-cine/pkg/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type MoviesController struct {
	conn *pgxpool.Pool
}

func (mc *MoviesController) List() []models.Movie {
	return []models.Movie{}
}

func (mc *MoviesController) Create(movie models.NewMovie) error {
	return nil
}

func (mc *MoviesController) Edit(models.Movie) error {
	return nil
}

func (mc *MoviesController) Delete(id int) error {
	return nil
}
